CREATE SEQUENCE IF NOT EXISTS order_number_seq;

CREATE OR REPLACE FUNCTION generate_order_number() 
RETURNS varchar(50) AS $$
BEGIN
    RETURN 'ORD-' || to_char(now(), 'YYYYMMDD') || '-' || lpad(nextval('order_number_seq')::text, 6, '0');
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION public.create_order_from_cart(p_user_id bigint, p_cart_store_id uuid, p_payment_method_id bigint, p_shipping_method_id bigint, p_shipping_addresses_id uuid, p_notes text DEFAULT NULL::text)
 RETURNS bigint
 LANGUAGE plpgsql
AS $function$
DECLARE
    v_order_id bigint;
    v_shipping_cost float8;
    v_total_price float8 := 0;
    v_final_price float8;
    v_order_number varchar(50);
    v_cart_item record;
    v_cart_id bigint;
BEGIN
    -- Dapatkan cart_id dan verifikasi kepemilikan user
    SELECT cs.cart_id INTO v_cart_id 
    FROM cart_stores cs
    JOIN carts c ON cs.cart_id = c.id
    WHERE cs.id = p_cart_store_id AND c.user_id = p_user_id;
    
    IF v_cart_id IS NULL THEN
        RAISE EXCEPTION 'Cart store dengan ID % tidak ditemukan atau bukan milik user %', p_cart_store_id, p_user_id;
    END IF;
    
    -- Get shipping cost
    SELECT price INTO v_shipping_cost FROM shipping_methods WHERE id = p_shipping_method_id;
    IF v_shipping_cost IS NULL THEN
        RAISE EXCEPTION 'Invalid shipping method ID %', p_shipping_method_id;
    END IF;
    
    -- Generate order number
    v_order_number := generate_order_number();
    
    -- Create order
    INSERT INTO orders (
        user_id,
        order_number,
        status_id,
        payment_method_id,
        shipping_method_id,
        shipping_addresses_id,
        shipping_cost,
        total_price,
        final_price,
        notes
    ) VALUES (
        p_user_id,
        v_order_number,
        1, -- Pending status
        p_payment_method_id,
        p_shipping_method_id,
        p_shipping_addresses_id,
        v_shipping_cost,
        0, -- Will be calculated
        0, -- Will be calculated
        p_notes
    ) RETURNING id INTO v_order_id;
    
    -- Process cart items
    FOR v_cart_item IN SELECT * FROM cart_items WHERE cart_store_id = p_cart_store_id
    LOOP
        -- Add order item
        INSERT INTO order_items (
            order_id,
            product_id,
            toko_id,
            quantity,
            price,
            discount_price,
            discount,
            subtotal
        ) VALUES (
            v_order_id,
            v_cart_item.product_id,
            (SELECT toko_id FROM cart_stores WHERE id = p_cart_store_id),
            v_cart_item.quantity,
            (SELECT price FROM products WHERE id = v_cart_item.product_id),
            (SELECT discount_price FROM products WHERE id = v_cart_item.product_id),
            (SELECT discount FROM products WHERE id = v_cart_item.product_id),
            (SELECT price FROM products WHERE id = v_cart_item.product_id) * v_cart_item.quantity
        );
        
        -- Update total price
        v_total_price := v_total_price + (SELECT price FROM products WHERE id = v_cart_item.product_id) * v_cart_item.quantity;
        
        -- Update product stock and sold count
        UPDATE products 
        SET stock = stock - v_cart_item.quantity, 
            sold = sold + v_cart_item.quantity,
            updated_at = now()
        WHERE id = v_cart_item.product_id;
    END LOOP;
    
    -- Calculate final price (total + shipping)
    v_final_price := v_total_price + v_shipping_cost;
    
    -- Update order with calculated prices
    UPDATE orders 
    SET total_price = v_total_price,
        final_price = v_final_price,
        updated_at = now()
    WHERE id = v_order_id;
    
    -- Add initial order tracking
    INSERT INTO order_tracking (order_id, status_id, notes)
    VALUES (v_order_id, 1, 'Order created');
    
    -- Clear the cart
    DELETE FROM cart_items WHERE cart_store_id = p_cart_store_id;
    DELETE FROM cart_stores WHERE id = p_cart_store_id;
    
    RETURN v_order_id;
END;
$function$
;


-- Function to update order status
CREATE OR REPLACE FUNCTION update_order_status(
    p_order_id bigint,
    p_status_id bigint,
    p_notes text DEFAULT NULL
) RETURNS void AS $$
BEGIN
    -- Update order status
    UPDATE orders 
    SET status_id = p_status_id,
        updated_at = now()
    WHERE id = p_order_id;
    
    -- Add tracking record
    INSERT INTO order_tracking (order_id, status_id, notes)
    VALUES (p_order_id, p_status_id, p_notes);
END;
$$ LANGUAGE plpgsql;
