import json
import re
import random


def convert_price_to_number(file_path):
    # Baca konten file
    with open(file_path, "r", encoding="utf-8") as file:
        content = file.read()

    # Ekstrak string JSON dari file Go
    start_index = content.find("`[")
    end_index = content.rfind("]`") + 1
    json_str = content[start_index + 1 : end_index]

    # Parse JSON string menjadi list of dictionaries
    data = json.loads(json_str)

    # Inisialisasi ID increment
    current_id = 1

    # Ubah price dan discount_price dari string menjadi number
    for item in data:
        if "price" in item and item["price"] != "null":
            # Hapus "Rp" dan karakter non-digit lainnya (seperti titik sebagai pemisah ribuan)
            clean_price = re.sub(r"[^\d]", "", item["price"])
            item["price"] = int(clean_price)

        # Juga ubah discount_price jika ada
        if "discount_price" in item and item["discount_price"] != "null":
            clean_discount = re.sub(r"[^\d]", "", item["discount_price"])
            item["discount_price"] = int(clean_discount)

        # Tambahkan random quantity
        item["quantity"] = random.randint(1, 100)

        # Tambahkan ID increment
        item["id"] = current_id
        current_id += 1

    # Konversi kembali ke string JSON dengan format yang rapi
    updated_json_str = json.dumps(data, indent=2, ensure_ascii=False)

    # Buat konten file Go yang baru
    updated_content = (
        content[:start_index] + "`" + updated_json_str + "`" + content[end_index + 1 :]
    )

    # Tulis kembali ke file (atau ke file baru)
    output_path = file_path.replace(".go", "_updated.go")
    with open(output_path, "w", encoding="utf-8") as file:
        file.write(updated_content)

    print(f"File telah diperbarui dan disimpan di {output_path}")
    # Tampilkan beberapa contoh data yang telah diubah
    print("\nBeberapa contoh data yang telah dikonversi:")
    for i in range(min(3, len(data))):
        print(f"Item {i+1}: {data[i]['product_name']}")
        print(f"  Price: {data[i]['price']} (tipe: {type(data[i]['price']).__name__})")
        print(
            f"  Discount Price: {data[i]['discount_price']} (tipe: {type(data[i]['discount_price']).__name__ if data[i]['discount_price'] is not None else 'None'})"
        )
        print(f"  Quantity: {data[i]['quantity']}")
        print(f"  ID: {data[i]['id']}")


# Panggil fungsi dengan path file
convert_price_to_number(
    "/media/jagres/HdJagres/backup/GoLang/api-tokopedia/internal/db/dummy/otomotif.go"
)
