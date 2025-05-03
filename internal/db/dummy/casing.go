package dummy

type Product struct {
	ProductName   string   `json:"product_name"`
	Price         string   `json:"price"`
	Stock         string   `json:"stock"`
	Description   string   `json:"description"`
	ImageURL      []string `json:"image_url"`
	Category      string   `json:"category"`
	Slug          string   `json:"slug"`
	Discount      string   `json:"discount"`
	DiscountPrice string   `json:"discount_price"`
	Rating        string   `json:"rating"`
	Sold          string   `json:"sold"`
	Country       string   `json:"country"`
	Estimation    string   `json:"estimation"`
	Toko          Toko     `json:"toko"`
}

type Toko struct {
	Username  string `json:"username"`
	TokoName  string `json:"toko_name"`
	ImageToko string `json:"image_toko"`
}

var CasingData = `[
  {
    "product_name": "Sticker Laptop Casing A Case Cover Hp 840 G1",
    "price": "Rp35.000",
    "stock": "17",
    "description": "Sticker Laptop Casing A Case Cover Hp 840 G1 ",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/12/8/26622886-b059-4289-8495-e78a512f9a36.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/12/8/0ecff8ae-24c8-477c-b339-17164707ee79.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "sticker-laptop-casing-a-case-cover-hp-840-g1",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 18",
    "country": "Jakarta Pusat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "slccom",
      "toko_name": "slccomputer",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/6/23/883e39fc-303c-44af-8fc4-ecbb58ff4bdf.jpg"
    }
  },
  {
    "product_name": "Sticker Laptop HP 830 G5 Silver Case Casing A Cover A HP 830 G5 Silver sparepart laptop sparepart laptop",
    "price": "Rp35.000",
    "stock": "19",
    "description": "Sticker Laptop Case Casing A Cover A HP 830 G5 Silver",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/10/14/19dba497-618c-4fa2-bd5a-14292d438b96.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/10/14/dbe1ce46-3579-4f91-ae3e-2de979ca59f7.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "sticker-laptop-hp-830-g5-silver-case-casing-a-cover-a-hp-830-g5-silver-sparepart-laptop-sparepart-laptop",
    "discount": "null",
    "discount_price": "null",
    "rating": "null",
    "sold": "null",
    "country": "Jakarta Pusat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "slccom",
      "toko_name": "slccomputer",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/6/23/883e39fc-303c-44af-8fc4-ecbb58ff4bdf.jpg"
    }
  },
  {
    "product_name": "Patrick Eye Case-Casing HP/Softcase/Hardcase for iPhone 11 12 13 14 15 - 11 Promax",
    "price": "Rp35.000",
    "stock": "8",
    "description": "✨ Happiness starts from little things ✨<br>Kelly menyeleksi berbagai aksesoris gemas dengan kualitas terjamin, harga terjangkau, dan tentunya layanan terbaik hanya untuk kamu. 💕<br>------------------------------------------------------<br>🔆NOTE PENTING : GANTUNGAN CASE DIJUAL TERPISAH KLIK VARIAN GANTUNGAN<br><br>Spesifikasi:<br>1. Berat : 42 gr (1 kg isi 30 pcs)<br>2. Jenis : hard case<br>3. Bahan : acrylic<br>4. Dimensi Kemasan : 23,2 x 10 x 1 cm<br>5. Kemasan : plastik OPP dengan seal lem penutup<br>6. Model : 1 warna (CIIH-001)<br>------------------------------------------------------<br><br>5 ⭐EXPRESS DELIVERY - Pengiriman setiap hari Senin sd Minggu. Libur nasional tetap proses order.<br>- kurir reguler pembayaran sebelum 16.00 WIB<br>- kurir sameday &amp; instant sebelum 14.30 WIB.<br><br>⭐ 7 HARI GARANSI KEPUASAN KU - Garansi 7 hari untuk pengembalian &amp; penukaran barang dengan free ongkir . Jaminan 100% puas.<br><br>⭐CS FAST-RESPONSE<br><br>⭐READY STOCK - Produk ready stock siap kirim dari Surabaya, bukan luar negeri dan bukan pre-order.",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/4/5/bdeb4893-3c8a-4e50-bc19-f71200dadfe0.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/4/5/50f65063-6d40-4a15-af0a-fa6fd43024d0.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/4/5/4283976f-8116-44c5-addd-669690de6d5d.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/4/5/692a6062-c8ed-41a6-a4db-84fee9320f5c.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/4/5/75e70767-0c7d-4f23-a971-e724ded6ec3c.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "patrick-eye-case-casing-hp-softcase-hardcase-for-iphone-11-12-13-14-15-11-promax-8426b",
    "discount": "50%",
    "discount_price": "Rp70.000",
    "rating": "5",
    "sold": "Terjual 1",
    "country": "Kota Surabaya",
    "estimation": "Estimasi tiba 31 Okt - 4 Nov",
    "toko": {
      "username": "minerhomestore",
      "toko_name": "Minerhome Store",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/11/27/bc48d792-dff9-4a08-827d-ff4df50aa7a5.png"
    }
  },
  {
    "product_name": "Tas HP kalibre SmartPhone Case 928069000",
    "price": "Rp189.000",
    "stock": "29",
    "description": "Kalibre Smartphone Case  adalah tas penyimpanan gadget, smartphone dan barang-barang kecil laiinya dengan 2 kompartment utama ,ada partisi dalam, adjustable strap, bag padded<br>dimensi 20cmX10cmX6cm",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/1/6/c1c81537-0424-4441-9f18-8a6b973c5a99.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/1/6/6be3342b-8bfb-4dda-a154-d0d9d840d916.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/1/6/e6fa5401-0b9f-4cc8-8b6a-39cb506a7322.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "tas-hp-kalibre-smartphone-case-928069000",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.9",
    "sold": "Terjual 250+",
    "country": "Kota Bandung",
    "estimation": "Estimasi tiba besok - 4 Nov",
    "toko": {
      "username": "kalibreid",
      "toko_name": "Kalibre Official Shop",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2017/9/25/2490772/2490772_85137367-8f85-486a-ac3a-8a4a3614710a.jpg"
    }
  },
  {
    "product_name": "BASEUS Lets Go Airbag Waterproof Phone Case Casing Sarung HP Anti Air - YELLOW GREY",
    "price": "Rp99.000",
    "stock": "19",
    "description": "PERHATIAN : Ada 2 pilihan warna di iklan ini, dikirimkan sesuai dengan pilihan variant ya.. No komplain<br>-Tali Kuning (YELLOW GREY)<br>-Tali Hitam (BLACK GREY)<br><br>IPX8 Certified Waterproof, Responsive Touch Screen<br><br>Material : PVC + ABC + Silicone<br>Technology : Hot pressing<br>Applicable models : Phones under 7.2inch<br>",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2021/4/27/f013a912-7cda-4ec2-abf2-7e8ed56d7846.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2020/10/9/f01dd434-00e3-4261-925d-43d0ed07dd49.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2020/10/9/6e6616e8-6169-49bf-8d05-e9c0fe7bfb8f.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2020/10/9/cb559eef-fd57-4ad5-88cf-664a562c9473.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2020/10/9/d98f690a-b0f8-4408-971b-4766b15aeccd.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "baseus-lets-go-airbag-waterproof-phone-case-casing-sarung-hp-anti-air-yellow-grey",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 3 rb+",
    "country": "Jakarta Utara",
    "estimation": "Estimasi tiba besok - 2 Nov",
    "toko": {
      "username": "himtech",
      "toko_name": "HimTech",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2019/3/12/1012361/1012361_0f2add41-1b54-45b1-8bb5-8081b703545c.png"
    }
  },
  {
    "product_name": "IDEA CASE - SAMSUNG S24 FE Soft Case Silikon Clear EXtra Bening Ring Magsave Casing Handphone - CLEAR, SAMSUNG S24 FE",
    "price": "Rp25.000",
    "stock": "48",
    "description": "PERHATIAN !!!! DIBACA SEBELUM MENILAI PRODUK !! <br><br>CASE INI ADA PELAPIS PLASTIK BAGIAN LUAR DAN DALAM . SEBELUM MERESPON CASE LECET / BEKAS , harap dicopot dulu plastiknya 2 sisi ya.<br><br>1. Di Wajibkan Setiap Pembukaan Packing Melakukan Video , Jika Barang Rusak/Kekurangan Barang Akan Di Validasi Melalui Video , Jika Tidak Ada Bukti Video Maka Di Anggap Tidak Sah Untuk Komplain Maupun Resolusi<br><br>2. PERHATIAN: Kami Tidak Menerima Komplain Pada Saat Barang Sampai Di Anda Rusak / Pecah / Kerusakan Pada Saat Pengiriman. Jadi Mohon Kerja Samanya Untuk Membeli Kardus Pengaman.<br><br>3.Setiap pengiriman kami sudah dilakukan qc sebelum mempacking barang . <br><br>PAHAMI KUALITAS PRODUK DENGAN BAIK SEBELUM MEMBELI<br>* Ketersediaan stok bisa di tanyakan via chat atau diskusi<br>*Kami dengan senang hati melayani grosiran dan eceran<br><br>Masukkan:<br>1) Tanggapan/Ulasan Anda sangat penting, masukan positif sangat dihargai.<br>2) Jangan ragu untuk memberitahu kami tahu sebelum meninggalkan umpan balik negatif, kami akan mencoba yang terbaik untuk memuaskan anda kebutuhan.<br><br>\"Awas Tertipu Dengan Barang Sama Tapi Kualitasnya Buruk,Produk Yang Kami Jual Sudah Melewati Quality Check Yang Mana Pada Saat Sampai Di Tangan Anda Dalam Keadaan OK Dan Tidak Cacat\"<br><br>Kelebihan Case Berbahan TPU Dibandingkan Dengan Case Silicone Adalah Case TPU umumnya tidak melorot atau menjadi longgar dari waktu ke waktu, tidak mudah terkelupas, dan sangat tahan terhadap gesekan, goresan, tekanan dan benturan sehingga mampu melindungi gadget Anda dengan sangat baik<br><br>Keunggulan Produk :<br>- Bahan TPU Premium Bertekstur Lembut<br>- Nyaman di Genggam Tidak Licin<br>- Melindungi Dari Benturan Dengan Sempurna<br>- Anti Scratch (Tidak Mudah Tergores)<br>- Tidak Menganggu Sinyal Ponsel<br>- Tidak menimbulkan baret pada bodi ponsel<br>- Anti Fingerprint apabila jari kita berminyak atau kotor walaupun terkena Case akan sangat mudah untuk dibersihkan<br>- ANTI KUNING <br>- TEBAL 0.9 MM <br><br>Q: Kak ko ditoko lain harganya lebih Murah?<br><br>A: Memang Benar Kak Di Toko Lain Lebih Murah,Tapi Ya Itu Kak Biasanya Produk Murah Blm Tentu Produknya Bagus Takutnya Mengecewakan Kakak,Lebih Baik Sedikit Mahal Tapi Memuaskan Di Bandingkan Murah Tapi Mengecewakan<br><br>Case Bening cocok utk Cover untuk ponsel anda.<br><br>Keunggulan:<br>- Bagus untuk semua warna dasar case anda<br>- Texture Halus <br>- Melepas panas lebih baik daripada case umumnya<br>- Tetap jauh lebih slim daripada pakai case biasa<br>- Pemasangan cukup mudah <br>- Melindungi dari goresan ketika ditaruh di meja<br><br>VARIAN : <br>CLEAR <br>CLEAR MAGSAVE ( WIRELESS CHARGER ) <br><br>SILAKAN PILIH SESUAI VARIAN <br><br>THANKS <br><br>HAPPY SHOPPING <br><br>GOD BLESS YOU ALL",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/10/1/195b762b-f8a7-4046-bac1-73e52f5b21a1.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/10/1/8bec67b1-ded6-4123-be7a-4aaf2b94193f.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/10/1/80f86fd1-048b-435f-9db2-1345c42a102e.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/10/1/774a813d-55dd-477b-a5ae-dde3b4b05994.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/10/1/3fd71bc0-346d-4444-b135-7bc813488149.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "idea-case-samsung-s24-fe-soft-case-silikon-clear-extra-bening-ring-magsave-casing-handphone-clear-samsung-s24-fe-7712b",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.5",
    "sold": "Terjual 14",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "pamola",
      "toko_name": "PAMOLLA",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/3/6/a2c043a9-b0e5-4bbe-b1f5-86e7096e2443.jpg"
    }
  },
  {
    "product_name": "Oxford Cloth Pouch Bag Phone Case - Free Carabiner Keychain Cover Fabric Sleeve Casing Belt Card Holder Money Slot Max 6.9 inch Screen Compatible 2 HandPhone HP Gantungan Tas Ikat Pinggang Kartu Uang",
    "price": "Rp108.900",
    "stock": "99",
    "description": "Oxford Cloth Pouch Bag Phone Case - Cover Casing Belt Card Holder Slot <br><br>Informasi Penting<br><br>Mau Dropship? Langsung order aja, atau bisa hubungi kami<br><br>Minat? Langsung order aja, dengan catatan sebagai berikut :<br><br>- Pastikan tipe pesanan anda sudah sesuai, sesuaikan dengan judul produknya, jika ragu dengan judul dan gambar yang tertera silahkan tanyakan. Kami tidak menerima komplain penerimaan pesanan, jika produk yang anda pesan salah tipe. Mohon perhatikan judul dengan teliti. Jika bentuk pada gambar berbeda dengan produk yang anda inginkan, baiknya tanyakan.<br><br>- Warna produk dijual sesuai dengan foto, jika pada foto hanya tertera 1 warna, maka produk hanya memiliki 1 varian warna. Jika warna produk bervariasi, sesuaikan dengan pilihan varian aktif. Jika warna produk bervariasi, dan tidak memiliki pilihan varian produk, maka dapat ditanyakan kepada kami.<br><br>- Setiap pemesanan tanpa bertanya warna tersedia terlebih dahulu (Khusus produk dengan variasi warna pada gambar dan tidak memiliki pilihan varian produk) akan kami kirimkan produk dengan warna secara acak, jika warna yang dipilih tidak tersedia<br><br>- Semua produk kami tersedia, silahkan langsung lakukan pemesanan, jika produk masih aktif dan dapat langsung dibeli<br><br>- Pastikan tipe dan warna sudah sesuai dengan yang di inginkan sesuai keterangan judul dan gambar<br><br>- Jaminan produk berkualitas dari Official Store<br><br>- Jika masih ada yang kurang dimengerti atau perlu ditanyakan, langsung tanyakan pada kolom diskusi atau pesan, dalam waktu singkat kami akan segera membalas setiap pertanyaan anda<br><br>- Dengan senang hati kami akan selalu membantu pembeli dalam setiap pertanyaan detail, bahkan foto langsung produk. Tidak perlu sungkan bertanya<br><br>- Pastikan membaca informasi penting pada catatan toko kami",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/1/10/6b1c3f73-e5f2-4bd9-ad2f-393d75bea68b.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/1/10/55432b50-67c7-40e0-bc82-5ae8aa858903.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/1/10/51822d22-9584-46a4-869d-29b8895a7a63.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/1/10/1dfd02b2-e649-42aa-a17e-53154fe1a682.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/1/10/01b6d80a-751f-4c71-81a1-ff74982882f0.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "oxford-cloth-pouch-bag-phone-case-free-carabiner-keychain-cover-fabric-sleeve-casing-belt-card-holder-money-slot-max-6-9-inch-screen-compatible-2-handphone-hp-gantungan-tas-ikat-pinggang-kartu-uang",
    "discount": "1%",
    "discount_price": "Rp110.000",
    "rating": "4.9",
    "sold": "Terjual 60+",
    "country": "Kota Bandung",
    "estimation": "Estimasi tiba besok - 4 Nov",
    "toko": {
      "username": "logaypektay",
      "toko_name": "Logay Accessories",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2014/3/4/92398/92398_5e514210-a389-11e3-ac7c-14a02523fab8.jpg"
    }
  },
  {
    "product_name": "Free Palestine (3) - Case Hp Kaligrafi Nama (Tersedia Semua Tipe HP) - Ultimatte Black",
    "price": "Rp109.000",
    "stock": "99",
    "description": "AhviSunday&nbsp;adalah&nbsp;brand&nbsp;Custom&nbsp;Kaligrafi&nbsp;terlaris&nbsp;Di&nbsp;Indonesia.&nbsp;Ribuan&nbsp;pesanan&nbsp;untuk&nbsp;case&nbsp;HP&nbsp;sudah&nbsp;dipercayakan&nbsp;ke&nbsp;AhviSunday.<br><br>🔥&nbsp;Tersedia&nbsp;semua&nbsp;tipe&nbsp;HP&nbsp;:&nbsp;Vivo,&nbsp;Oppo,&nbsp;Xiaomi,&nbsp;Lenovo,&nbsp;Iphone,&nbsp;Samsung&nbsp;dll<br>🔥&nbsp;Kualitas&nbsp;Mesin&nbsp;UV&nbsp;asli&nbsp;Jepang<br>🔥&nbsp;Hasil&nbsp;cetak&nbsp;jelas&nbsp;dan&nbsp;tidak&nbsp;luntur<br>🔥&nbsp;Bahan&nbsp;ringan&nbsp;dalam&nbsp;genggaman<br>🔥&nbsp;Finishing&nbsp;Doff&nbsp;/&nbsp;Glossy<br><br>==================================================<br><br>Informasi&nbsp;Bahan&nbsp;:<br><br>Ultimatte&nbsp;Black<br>✔️&nbsp;Bahan&nbsp;dasar&nbsp;silikon<br>✔️&nbsp;Cetak&nbsp;bagian&nbsp;permukaan&nbsp;belakang<br>✔️&nbsp;Lentur&nbsp;dan&nbsp;fit&nbsp;dengan&nbsp;smartphone<br>✔️&nbsp;Dasar&nbsp;Case&nbsp;berwarna&nbsp;hitam<br><br>Clear&nbsp;Case<br>✔️&nbsp;Bahan&nbsp;dasar&nbsp;silikon&nbsp;bening<br>✔️&nbsp;Cetak&nbsp;bagian&nbsp;permukaan&nbsp;belakang<br>✔️&nbsp;Lentur&nbsp;dan&nbsp;fit&nbsp;dengan&nbsp;smartphone<br>✔️&nbsp;Elastis&nbsp;bisa&nbsp;kembali&nbsp;ke&nbsp;bentuk&nbsp;semula<br>✔️&nbsp;Ada&nbsp;sisi&nbsp;bumper&nbsp;di&nbsp;sudut&nbsp;untuk&nbsp;meminimalisir&nbsp;retak&nbsp;apabila&nbsp;terjatuh<br><br>Slim&nbsp;Clear&nbsp;Case<br>✔️&nbsp;Bahan&nbsp;dasar&nbsp;silikon&nbsp;bening<br>✔️&nbsp;Cetak&nbsp;bagian&nbsp;permukaan&nbsp;belakang<br>✔️&nbsp;Lentur&nbsp;dan&nbsp;fit&nbsp;dengan&nbsp;smartphone<br>✔️&nbsp;Elastis&nbsp;bisa&nbsp;kembali&nbsp;ke&nbsp;bentuk&nbsp;semula<br><br>Exclusive&nbsp;Case<br>✔️&nbsp;Anti&nbsp;Gores,&nbsp;Anti&nbsp;Api&nbsp;dan&nbsp;Anti&nbsp;Noda<br>✔️&nbsp;Full&nbsp;Glass&nbsp;/&nbsp;plat&nbsp;besi&nbsp;dengan&nbsp;Coating&nbsp;TPU<br>✔️&nbsp;Print&nbsp;pada&nbsp;permukaan<br>✔️&nbsp;Menggunakan&nbsp;sistem&nbsp;cetak&nbsp;Transfer&nbsp;Printing&nbsp;Film<br>✔️&nbsp;Bevel&nbsp;di&nbsp;ujung&nbsp;case&nbsp;agar&nbsp;glass&nbsp;tidak&nbsp;tajam<br><br>Fullprint&nbsp;Hardcase<br>✔️&nbsp;Bahan&nbsp;Polycarbonate<br>✔️&nbsp;Full&nbsp;print&nbsp;/&nbsp;cetak&nbsp;hingga&nbsp;sisi&nbsp;samping&nbsp;case<br>✔️&nbsp;Hasil&nbsp;jelas&nbsp;dan&nbsp;tidak&nbsp;luntur<br>✔️&nbsp;Atas&nbsp;dan&nbsp;bawah&nbsp;case&nbsp;terbuka<br>✔️&nbsp;Bahan&nbsp;ringan&nbsp;di&nbsp;dalam&nbsp;genggaman&nbsp;walau&nbsp;bahan&nbsp;keras<br><br>==================================================<br><br>Cara&nbsp;Pemesanan&nbsp;:<br><br>⚠️&nbsp;PENTING!!!&nbsp;Tanyakan&nbsp;ketersediaan&nbsp;stok&nbsp;case&nbsp;terlebih&nbsp;dahulu<br>✔️&nbsp;Pilih&nbsp;jenis&nbsp;case<br>✔️&nbsp;Masukkan&nbsp;gambar&nbsp;desain&nbsp;pada&nbsp;keterangan&nbsp;/&nbsp;chat&nbsp;via&nbsp;tokopedia<br>✔️&nbsp;Mohon&nbsp;cantumkan&nbsp;nama&nbsp;kaligrafi&nbsp;beserta&nbsp;tipe&nbsp;hpnya&nbsp;dicatatan&nbsp;pembelian<br><br>FREE&nbsp;Garansi&nbsp;/&nbsp;Retur&nbsp;Jika&nbsp;diterima&nbsp;kondisi&nbsp;rusak",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/08e0b5b3-91ab-4275-9469-639050a1c89c.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/11e0d8c4-8204-4bd2-bba8-8a8f79fc0948.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/0632d7f6-696a-4d88-900b-a714c7ecc730.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/08e0b5b3-91ab-4275-9469-639050a1c89c.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/a82f26c0-8486-4b4a-95e2-5503ab4a0e2f.png.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "free-palestine-3-case-hp-kaligrafi-nama-tersedia-semua-tipe-hp-ultimatte-black-bc8ee",
    "discount": "null",
    "discount_price": "null",
    "rating": "null",
    "sold": "null",
    "country": "Kab. Cianjur",
    "estimation": "Estimasi tiba 14 - 17 Nov",
    "toko": {
      "username": "ahvisunday",
      "toko_name": "AhviSunday",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2022/6/7/a74cd2c3-c2fd-40fa-a651-72321b1b6bf4.jpg"
    }
  },
  {
    "product_name": "AUKEY iPhone 15 Premium Silicone Case PC-GJ10A with MagSafe Casing Hp - Blue",
    "price": "Rp155.000",
    "stock": "8",
    "description": "AUKEY iPhone 15 Premium Silicone Case PC-GJ10A with MagSafe<br><br>HIGH QUALITY MATERIAL<br>Berbahan Medical-grade Liquid Silicone (Silikon berkualitas tinggi dan murni) yang dapat memberikan perlindungan pada perangkat dengan sentuhan yang lembut dan halus hingga aman untuk kulit.<br><br>KEEP DUST AND SMUDGES AWAY<br>Memiliki lapisan anti-statis yang mampu mencegah debu agar tidak menempel pada casing<br><br>PRECISE CUTOUTS<br>Cutout yang presisi yang ditunjukan dengan Ports dan tombol-tombol yang bisa dapat diakses sepenuhnya<br><br>COMPATIBLE WITH MAGSAFE<br>Memiliki magnet N52 bawaan yang kuat yang mendukung semua aksesoris maupun perangkat yang kompatibel dengan Magsafe<br><br>SCREEN AND CAMERA PROTECTION<br>Memberikan perlindungan sepenuhnya pada Kamera belakang dan layer handphone.<br><br>Spesifikasi :<br>Model: PC-GJ10A<br>Material: Silicone<br>Compatible for: iPhone 15<br>Magsafe : YES<br>Color: Midnight, Blue, Lilac, Succulent",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/11/25/7d253bd3-2a49-4167-8ee3-8a299430b8bf.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/11/25/8f7e4f6b-af61-4b60-b8b4-c49702c91119.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/11/25/fde90e1f-1576-46dd-8e56-0257259b2ee1.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/11/25/d2174191-0720-4fb1-a47f-1abcf2d3be16.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/11/25/0a815394-d3eb-44b9-a822-53efc6fd3689.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "aukey-iphone-15-premium-silicone-case-pc-gj10a-with-magsafe-casing-hp-blue-1f0ec",
    "discount": "48%",
    "discount_price": "Rp299.000",
    "rating": "4.9",
    "sold": "Terjual 250+",
    "country": "Jakarta Timur",
    "estimation": "Estimasi tiba 31 Okt - 2 Nov",
    "toko": {
      "username": "aukeyofficialstore",
      "toko_name": "AUKEY OFFICIAL STORE",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/9/12/a98899e5-8ed1-4728-991b-df14d2a409bb.jpg"
    }
  },
  {
    "product_name": "Casing Samsung Galaxy A35 A55 5G Silicone Cover Softcase Case Hp - Galaxy A35, Lime",
    "price": "Rp207.000",
    "stock": "7",
    "description": "BARANG ORIGINAL !!!<br><br>CHECK JUGA GADGET FAVORITMU DISINI<br>Smarphone : <a target=\"_blank\" rel=\"noopener noreferrer\" href=\"https://www.tokopedia.com/supergadgettt/etalase/smartphone\">https://www.tokopedia.com/supergadgettt/etalase/smartphone</a><br>Tablet : <a target=\"_blank\" rel=\"noopener noreferrer\" href=\"https://www.tokopedia.com/supergadgettt/etalase/tablets\">https://www.tokopedia.com/supergadgettt/etalase/tablets</a><br><br><br>Samsung Galaxy Silicone Cover<br><br>Desain ramping<br>Perasaan sentuhan lembut<br>Meningkatkan kelengkapan<br>Bahan ramah lingkungan<br><br>Warna lembut, perlindungan keras<br>Casing Silikon melengkapi perangkat Anda dengan warna-warna pastel modis yang menambahkan lapisan perlindungan silikon yang kokoh<br><br>Pas sehalus sutra<br>Sehalus sutra dan mulus, Casing Silikon adalah lapisan padat yang dirancang untuk melindungi perangkat Anda dari goresan dan jatuh.<br><br>Sentuhan yang menyenangkan, pegangan yang aman<br>Casing Silikon terpasang dengan nyaman di tangan Anda dan digenggam erat di telapak tangan Anda.<br><br>Check Device Favorite disini :<br><a target=\"_blank\" rel=\"noopener noreferrer\" href=\"https://www.tokopedia.com/supergadgettt/etalase/samsung\">https://www.tokopedia.com/supergadgettt/etalase/samsung</a><br><br>Compatible models for : <br>Samsung Galaxy A35<br>Samsung Galaxy A55<br><br>Weight<br>43 g<br><br><br>Notes<br>- Pastikan sudah memilih varian model sudah tepat dan sesuai dengan model hp anda, karena judul produk dibuat secara universal untuk semua tipe hp.<br><br>** Perhatian Varian Model / Warna Produk Yang Dipilih Sebelum Melakukan Pembelian",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/3/12/8140af65-efea-4514-bf9c-1faa082790f1.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/3/12/5f8c24ed-e63e-40ec-8769-7b8a891264f0.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/3/12/16716d67-a587-447f-9e37-62632b3de791.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/3/12/95b46e21-d8e4-473d-8c52-51e2a8c92248.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/3/12/679ac09a-eb00-449b-9598-81db497c3c0e.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-samsung-galaxy-a35-a55-5g-silicone-cover-softcase-case-hp-galaxy-a35-lime-f763b",
    "discount": "10%",
    "discount_price": "Rp229.000",
    "rating": "4.9",
    "sold": "Terjual 100+",
    "country": "Jakarta Pusat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "supergadgettt",
      "toko_name": "SUPER_GADGET",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2019/8/8/405517/405517_a23f3177-3e6a-44cb-a01f-ff132e8cea0c.png"
    }
  },
  {
    "product_name": "Casing Samsung Galaxy Z Fold 6 Silicone Case Softcase Pelindung HP Z Fold6 - White",
    "price": "Rp749.000",
    "stock": "2",
    "description": "Fold6 Silicone Case<br><br>Dapat digunakan untuk Samsung Galaxy Z Fold6<br>Color: Gray, White<br><br>Dimension:<br>68 mm x 157 mm x 17<br><br>Weight:<br>0.04",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/11/bdf74e91-e180-46fb-b0c4-525227e73031.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/11/1b6c96de-951c-435b-a1b0-bc68e6ecb5e8.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/11/41844705-8deb-4698-9710-13c207e4ecbd.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/11/c20c33bf-a7ae-439e-94b0-86ea10f61dcf.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/11/3c5ebd7f-d348-4808-beb1-b6e96955b3b5.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/11/bdf74e91-e180-46fb-b0c4-525227e73031.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/13/dd1e82f0-22b0-4917-96e5-bd859319240a.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/11/1b6c96de-951c-435b-a1b0-bc68e6ecb5e8.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/11/41844705-8deb-4698-9710-13c207e4ecbd.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/11/c20c33bf-a7ae-439e-94b0-86ea10f61dcf.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/11/3c5ebd7f-d348-4808-beb1-b6e96955b3b5.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-samsung-galaxy-z-fold-6-silicone-case-softcase-pelindung-hp-z-fold6-white-f911f",
    "discount": "25%",
    "discount_price": "Rp999.000",
    "rating": "5",
    "sold": "Terjual 22",
    "country": "Jakarta Pusat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "supergadgettt",
      "toko_name": "SUPER_GADGET",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2019/8/8/405517/405517_a23f3177-3e6a-44cb-a01f-ff132e8cea0c.png"
    }
  },
  {
    "product_name": "Case Casing hp iPhone 6 6S 6+ 6S+ 7 7+ 8 8+ Plus SE Kartun Flowers - Hijau terang, iPhone 6",
    "price": "Rp25.000",
    "stock": "997",
    "description": "Selamat datang di toko kami ❤❤❤<br>Produk yang diperintahkan akan pabrik Indonesia dalam satu hari kerja, dan kami bergegas ke arahmu!<br><br>Kenapa kau memilih kami?<br>🤙KUALITAS TINGGI<br>Casing ponsel ini terbuat dari bahan Casing hp berkualitas tinggi serta, mudah dipasang &amp; dilepas, memberikan perlindungan terbaik untuk HP anda<br><br>🤙PENGIRIMAN LOKAL<br>Pengiriman langsung dari pabrik Indonesia untuk memastikan pengiriman barang yang cepat dan aman ke tangan Anda.<br><br>🤙LAYANAN YANG BAIK<br>Jika Anda memiliki pertanyaan tentang produk, layanan pelanggan akan meminta Anda online untuk layanan<br><br>Deskripsi Produk:<br>🤙💯Baru dan Dapat Dihapus.<br>🤙Desain pola kotak mode membuat ponsel Anda tetap pribadi dan terasa nyaman.<br>🤙Penyerap goncangan, bahan anti pecah.<br>🤙Lanyard menambahkan perlindungan terhadap jatuh yang tidak disengaja.<br>🤙Semua fungsi berfungsi dengan baik bahkan ponsel Anda dalam kasus ini.<br><br>Fitur:<br>Bahan: Terbuat dari Silikon Berkualitas Tinggi + TPU<br>Isi paket:<br>1 x Softcase elektroplating + 1 x Lanyard",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/6/20/bd78a773-c090-4ca9-b355-73620515c077.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/6/20/3632f15e-e66a-4483-8fa5-ab1e6cf3fe56.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/6/20/c6c74543-32f6-418b-9f2f-b29d40322bd2.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/6/20/363d39e5-4a87-4be5-927b-428543f39797.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/6/20/7381560d-ba6a-48b0-9642-dc10670a42be.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "case-casing-hp-iphone-6-6s-6-6s-7-7-8-8-plus-se-kartun-flowers-hijau-terang-iphone-6",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.9",
    "sold": "Terjual 40+",
    "country": "Kab. Tangerang",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "shopeper",
      "toko_name": "Trendi Case",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/4/29/1df532de-5202-4b3e-8db7-a42d37a779de.jpg"
    }
  },
  {
    "product_name": "Casing hp iPhone 6 6S 6+ 6S+ 7 7+ 8 8+ Plus SE(2020) Kartun Karakter - Gaya 1, iPhone 6",
    "price": "Rp0",
    "stock": "Habis",
    "description": "Selamat datang di toko kami ❤❤❤<br>Produk yang diperintahkan akan pabrik Indonesia dalam satu hari kerja, dan kami bergegas ke arahmu!<br><br>Kenapa kau memilih kami?<br>🤙KUALITAS TINGGI<br>Casing ponsel ini terbuat dari bahan softcase berkualitas tinggi serta, mudah dipasang &amp; dilepas, memberikan perlindungan terbaik untuk HP anda<br><br>🤙PENGIRIMAN LOKAL<br>Pengiriman langsung dari pabrik Indonesia untuk memastikan pengiriman barang yang cepat dan aman ke tangan Anda.<br><br>🤙LAYANAN YANG BAIK<br>Jika Anda memiliki pertanyaan tentang produk, layanan pelanggan akan meminta Anda online untuk layanan<br><br>Deskripsi Produk:<br>🤙💯Brand baru dan dapat dihapus.<br>🤙Desain wavy mode menjaga ponselmu pribadi dan merasa baik.<br>🤙Membebaskan pola yang sama hadiah pemilik kembali.<br>🤙Pemilik belakang mendukung lokasi telepon 360 derajat ketika digunakan.<br>🤙Penyerap kejutan, materi silikon lembut yang tak bisa hancur.<br>🤙Desain lensa unik untuk menghindari goresan kamera.<br>🤙Lubang kamera yang tebal menambah perlindungan terhadap titik kecelakaan.<br>🤙Bahkan jika ponselmu berada di kondisi ini, semua fungsi akan bekerja dengan baik.<br><br>Fitur:<br>Bahan: Dibuat dari TPU Kualitas Tinggi<br>Isi paket:<br>1 x Kasing",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/28/30390fef-2f93-4b7e-a85b-279d03bce1ca.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/28/c1cfda92-b7fc-46cc-91c6-3592cb6d6da8.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/28/3d32d1b5-0523-4765-8ba9-d69886db7294.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/28/daebd170-1ebc-4487-ad79-212f21b76465.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/28/a6ccaabb-c239-42b9-bd79-ef907c225de4.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-iphone-6-6s-6-6s-7-7-8-8-plus-se-2020-kartun-karakter-gaya-1-iphone-6",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 70+",
    "country": "null",
    "estimation": "null",
    "toko": {
      "username": "shopeper",
      "toko_name": "Trendi Case",
      "image_toko": ""
    }
  },
  {
    "product_name": "Madinah (Indoor) - Case Hp Kaligrafi Nama (Tersedia Semua Tipe HP) - Ultimatte Black",
    "price": "Rp109.000",
    "stock": "98",
    "description": "AhviSunday&nbsp;adalah&nbsp;brand&nbsp;Custom&nbsp;Kaligrafi&nbsp;terlaris&nbsp;Di&nbsp;Indonesia.&nbsp;Ribuan&nbsp;pesanan&nbsp;untuk&nbsp;case&nbsp;HP&nbsp;sudah&nbsp;dipercayakan&nbsp;ke&nbsp;AhviSunday.<br><br>🔥&nbsp;Tersedia&nbsp;semua&nbsp;tipe&nbsp;HP&nbsp;:&nbsp;Vivo,&nbsp;Oppo,&nbsp;Xiaomi,&nbsp;Lenovo,&nbsp;Iphone,&nbsp;Samsung&nbsp;dll<br>🔥&nbsp;Kualitas&nbsp;Mesin&nbsp;UV&nbsp;asli&nbsp;Jepang<br>🔥&nbsp;Hasil&nbsp;cetak&nbsp;jelas&nbsp;dan&nbsp;tidak&nbsp;luntur<br>🔥&nbsp;Bahan&nbsp;ringan&nbsp;dalam&nbsp;genggaman<br>🔥&nbsp;Finishing&nbsp;Doff&nbsp;/&nbsp;Glossy<br><br>==================================================<br><br>Informasi&nbsp;Bahan&nbsp;:<br><br>Ultimatte&nbsp;Black<br>✔️&nbsp;Bahan&nbsp;dasar&nbsp;silikon<br>✔️&nbsp;Cetak&nbsp;bagian&nbsp;permukaan&nbsp;belakang<br>✔️&nbsp;Lentur&nbsp;dan&nbsp;fit&nbsp;dengan&nbsp;smartphone<br>✔️&nbsp;Dasar&nbsp;Case&nbsp;berwarna&nbsp;hitam<br><br>Clear&nbsp;Case<br>✔️&nbsp;Bahan&nbsp;dasar&nbsp;silikon&nbsp;bening<br>✔️&nbsp;Cetak&nbsp;bagian&nbsp;permukaan&nbsp;belakang<br>✔️&nbsp;Lentur&nbsp;dan&nbsp;fit&nbsp;dengan&nbsp;smartphone<br>✔️&nbsp;Elastis&nbsp;bisa&nbsp;kembali&nbsp;ke&nbsp;bentuk&nbsp;semula<br>✔️&nbsp;Ada&nbsp;sisi&nbsp;bumper&nbsp;di&nbsp;sudut&nbsp;untuk&nbsp;meminimalisir&nbsp;retak&nbsp;apabila&nbsp;terjatuh<br><br>Slim&nbsp;Clear&nbsp;Case<br>✔️&nbsp;Bahan&nbsp;dasar&nbsp;silikon&nbsp;bening<br>✔️&nbsp;Cetak&nbsp;bagian&nbsp;permukaan&nbsp;belakang<br>✔️&nbsp;Lentur&nbsp;dan&nbsp;fit&nbsp;dengan&nbsp;smartphone<br>✔️&nbsp;Elastis&nbsp;bisa&nbsp;kembali&nbsp;ke&nbsp;bentuk&nbsp;semula<br><br>Exclusive&nbsp;Case<br>✔️&nbsp;Anti&nbsp;Gores,&nbsp;Anti&nbsp;Api&nbsp;dan&nbsp;Anti&nbsp;Noda<br>✔️&nbsp;Full&nbsp;Glass&nbsp;/&nbsp;plat&nbsp;besi&nbsp;dengan&nbsp;Coating&nbsp;TPU<br>✔️&nbsp;Print&nbsp;pada&nbsp;permukaan<br>✔️&nbsp;Menggunakan&nbsp;sistem&nbsp;cetak&nbsp;Transfer&nbsp;Printing&nbsp;Film<br>✔️&nbsp;Bevel&nbsp;di&nbsp;ujung&nbsp;case&nbsp;agar&nbsp;glass&nbsp;tidak&nbsp;tajam<br><br>Fullprint&nbsp;Hardcase<br>✔️&nbsp;Bahan&nbsp;Polycarbonate<br>✔️&nbsp;Full&nbsp;print&nbsp;/&nbsp;cetak&nbsp;hingga&nbsp;sisi&nbsp;samping&nbsp;case<br>✔️&nbsp;Hasil&nbsp;jelas&nbsp;dan&nbsp;tidak&nbsp;luntur<br>✔️&nbsp;Atas&nbsp;dan&nbsp;bawah&nbsp;case&nbsp;terbuka<br>✔️&nbsp;Bahan&nbsp;ringan&nbsp;di&nbsp;dalam&nbsp;genggaman&nbsp;walau&nbsp;bahan&nbsp;keras<br><br>==================================================<br><br>Cara&nbsp;Pemesanan&nbsp;:<br><br>⚠️&nbsp;PENTING!!!&nbsp;Tanyakan&nbsp;ketersediaan&nbsp;stok&nbsp;case&nbsp;terlebih&nbsp;dahulu<br>✔️&nbsp;Pilih&nbsp;jenis&nbsp;case<br>✔️&nbsp;Masukkan&nbsp;gambar&nbsp;desain&nbsp;pada&nbsp;keterangan&nbsp;/&nbsp;chat&nbsp;via&nbsp;tokopedia<br>✔️&nbsp;Mohon&nbsp;cantumkan&nbsp;nama&nbsp;kaligrafi&nbsp;beserta&nbsp;tipe&nbsp;hpnya&nbsp;dicatatan&nbsp;pembelian<br><br>FREE&nbsp;Garansi&nbsp;/&nbsp;Retur&nbsp;Jika&nbsp;diterima&nbsp;kondisi&nbsp;rusak",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/22/47f5a117-144f-4b11-832b-1f22d1c92e3b.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/22/126e2b93-9f40-4d89-b158-b04c34880f6e.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/22/5e75adf2-0e2f-4ad7-bbbc-d24adac4d74c.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/22/47f5a117-144f-4b11-832b-1f22d1c92e3b.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/22/6b7979f7-6442-4392-a2d9-bf328d84bd2f.png.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "madinah-indoor-case-hp-kaligrafi-nama-tersedia-semua-tipe-hp-ultimatte-black-2990f",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 6",
    "country": "Kab. Cianjur",
    "estimation": "Estimasi tiba 14 - 17 Nov",
    "toko": {
      "username": "ahvisunday",
      "toko_name": "AhviSunday",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2022/6/7/a74cd2c3-c2fd-40fa-a651-72321b1b6bf4.jpg"
    }
  },
  {
    "product_name": "Case hp iphone 11 / 11 pro / 11 pro max Hard case magnetic - Gold, Iphone 11",
    "price": "Rp75.000",
    "stock": "100",
    "description": "wellcome<br><br>casing&nbsp;hardcase&nbsp;magsafe&nbsp;pc<br><br>tersedia&nbsp;warna<br>~&nbsp;Black<br>~&nbsp;Gold<br>~&nbsp;Rose&nbsp;Gold<br>~&nbsp;Deep&nbsp;purple<br><br>happy&nbsp;shopping",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/1/21/475c5058-9920-4b4a-9ba9-1567dbd83fe5.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/1/21/efee54ad-523c-4367-b66b-3dfd5ddd11e3.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/1/21/77d9bead-9e67-4f8f-949e-743dea9ada6b.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/1/21/1211817e-ceee-443e-bc0d-91b7995e0413.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "case-hp-iphone-11-11-pro-11-pro-max-hard-case-magnetic-gold-iphone-11-2ff6d",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.8",
    "sold": "Terjual 100+",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 2 Nov",
    "toko": {
      "username": "delvinstore",
      "toko_name": "Delvin@store",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2024/1/21/c5e67e1d-e977-43d7-910a-da747b26edff.jpg"
    }
  },
  {
    "product_name": "REDMI A1 A2 A3 8 8A PRO 9 9A 9C 9T 10 10A 11 LITE 11T 12 12C 13 13C HOLOGRAM CASING VIRAL KESING HP CESING HANDPHONE SOFT HARD CASE XIAOMI - Hitam, REDMI 8",
    "price": "Rp15.000",
    "stock": "1.000",
    "description": "KELEBIHAN PRODUK :<br><br>Bahan material : Hybrid campuran softcase dan hardcase<br>- Bahan TPU Premium Bertekstur Lembut<br>- Nyaman di Genggam Tidak Licin<br>- Melindungi Dari Benturan Dengan Sempurna<br>- Anti Scratch (Tidak Mudah Tergores)<br>- Tidak Menganggu Sinyal Ponsel<br>- Tidak menimbulkan baret pada bodi ponsel<br>- Anti Fingerprint apabila jari kita berminyak atau kotor walaupun terkena Case akan sangat mudah untuk dibersihkan<br><br>MEMBELI = SETUJU<br>tidak menerima returan jika tidak ada video unboxing<br>#Liang's Store",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/0e7f1e1e-7cd3-4e60-8e33-e7c04ae86749.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/111f155b-d968-4439-9b6e-651fecb145f0.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/5023c398-8ccf-4321-a613-d82cc724661f.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/244711a7-e17f-428b-b114-c12ab6015cdd.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "redmi-a1-a2-a3-8-8a-pro-9-9a-9c-9t-10-10a-11-lite-11t-12-12c-13-13c-hologram-casing-viral-kesing-hp-cesing-handphone-soft-hard-case-xiaomi-hitam-redmi-8-b80ef",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 13",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "liangsstore",
      "toko_name": "Liang's Store",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2022/10/11/39000dc2-26cd-4750-988d-968d4c5b8bed.jpg"
    }
  },
  {
    "product_name": "Kakbah (Blue Sky) - Case Hp Kaligrafi Nama (Tersedia Semua Tipe HP) - Ultimatte Black",
    "price": "Rp109.000",
    "stock": "98",
    "description": "AhviSunday&nbsp;adalah&nbsp;brand&nbsp;Custom&nbsp;Kaligrafi&nbsp;terlaris&nbsp;Di&nbsp;Indonesia.&nbsp;Ribuan&nbsp;pesanan&nbsp;untuk&nbsp;case&nbsp;HP&nbsp;sudah&nbsp;dipercayakan&nbsp;ke&nbsp;AhviSunday.<br><br>🔥&nbsp;Tersedia&nbsp;semua&nbsp;tipe&nbsp;HP&nbsp;:&nbsp;Vivo,&nbsp;Oppo,&nbsp;Xiaomi,&nbsp;Lenovo,&nbsp;Iphone,&nbsp;Samsung&nbsp;dll<br>🔥&nbsp;Kualitas&nbsp;Mesin&nbsp;UV&nbsp;asli&nbsp;Jepang<br>🔥&nbsp;Hasil&nbsp;cetak&nbsp;jelas&nbsp;dan&nbsp;tidak&nbsp;luntur<br>🔥&nbsp;Bahan&nbsp;ringan&nbsp;dalam&nbsp;genggaman<br>🔥&nbsp;Finishing&nbsp;Doff&nbsp;/&nbsp;Glossy<br><br>==================================================<br><br>Informasi&nbsp;Bahan&nbsp;:<br><br>Ultimatte&nbsp;Black<br>✔️&nbsp;Bahan&nbsp;dasar&nbsp;silikon<br>✔️&nbsp;Cetak&nbsp;bagian&nbsp;permukaan&nbsp;belakang<br>✔️&nbsp;Lentur&nbsp;dan&nbsp;fit&nbsp;dengan&nbsp;smartphone<br>✔️&nbsp;Dasar&nbsp;Case&nbsp;berwarna&nbsp;hitam<br><br>Clear&nbsp;Case<br>✔️&nbsp;Bahan&nbsp;dasar&nbsp;silikon&nbsp;bening<br>✔️&nbsp;Cetak&nbsp;bagian&nbsp;permukaan&nbsp;belakang<br>✔️&nbsp;Lentur&nbsp;dan&nbsp;fit&nbsp;dengan&nbsp;smartphone<br>✔️&nbsp;Elastis&nbsp;bisa&nbsp;kembali&nbsp;ke&nbsp;bentuk&nbsp;semula<br>✔️&nbsp;Ada&nbsp;sisi&nbsp;bumper&nbsp;di&nbsp;sudut&nbsp;untuk&nbsp;meminimalisir&nbsp;retak&nbsp;apabila&nbsp;terjatuh<br><br>Slim&nbsp;Clear&nbsp;Case<br>✔️&nbsp;Bahan&nbsp;dasar&nbsp;silikon&nbsp;bening<br>✔️&nbsp;Cetak&nbsp;bagian&nbsp;permukaan&nbsp;belakang<br>✔️&nbsp;Lentur&nbsp;dan&nbsp;fit&nbsp;dengan&nbsp;smartphone<br>✔️&nbsp;Elastis&nbsp;bisa&nbsp;kembali&nbsp;ke&nbsp;bentuk&nbsp;semula<br><br>Exclusive&nbsp;Case<br>✔️&nbsp;Anti&nbsp;Gores,&nbsp;Anti&nbsp;Api&nbsp;dan&nbsp;Anti&nbsp;Noda<br>✔️&nbsp;Full&nbsp;Glass&nbsp;/&nbsp;plat&nbsp;besi&nbsp;dengan&nbsp;Coating&nbsp;TPU<br>✔️&nbsp;Print&nbsp;pada&nbsp;permukaan<br>✔️&nbsp;Menggunakan&nbsp;sistem&nbsp;cetak&nbsp;Transfer&nbsp;Printing&nbsp;Film<br>✔️&nbsp;Bevel&nbsp;di&nbsp;ujung&nbsp;case&nbsp;agar&nbsp;glass&nbsp;tidak&nbsp;tajam<br><br>Fullprint&nbsp;Hardcase<br>✔️&nbsp;Bahan&nbsp;Polycarbonate<br>✔️&nbsp;Full&nbsp;print&nbsp;/&nbsp;cetak&nbsp;hingga&nbsp;sisi&nbsp;samping&nbsp;case<br>✔️&nbsp;Hasil&nbsp;jelas&nbsp;dan&nbsp;tidak&nbsp;luntur<br>✔️&nbsp;Atas&nbsp;dan&nbsp;bawah&nbsp;case&nbsp;terbuka<br>✔️&nbsp;Bahan&nbsp;ringan&nbsp;di&nbsp;dalam&nbsp;genggaman&nbsp;walau&nbsp;bahan&nbsp;keras<br><br>==================================================<br><br>Cara&nbsp;Pemesanan&nbsp;:<br><br>⚠️&nbsp;PENTING!!!&nbsp;Tanyakan&nbsp;ketersediaan&nbsp;stok&nbsp;case&nbsp;terlebih&nbsp;dahulu<br>✔️&nbsp;Pilih&nbsp;jenis&nbsp;case<br>✔️&nbsp;Masukkan&nbsp;gambar&nbsp;desain&nbsp;pada&nbsp;keterangan&nbsp;/&nbsp;chat&nbsp;via&nbsp;tokopedia<br>✔️&nbsp;Mohon&nbsp;cantumkan&nbsp;nama&nbsp;kaligrafi&nbsp;beserta&nbsp;tipe&nbsp;hpnya&nbsp;dicatatan&nbsp;pembelian<br><br>FREE&nbsp;Garansi&nbsp;/&nbsp;Retur&nbsp;Jika&nbsp;diterima&nbsp;kondisi&nbsp;rusak",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/f787bad8-1b8a-44cd-a3bf-f5be89229e70.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/9c481c3f-b90a-4bc9-baed-10d91c2b5ee7.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/df584466-9a6b-4a7c-b0ae-ef98c48df714.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/f787bad8-1b8a-44cd-a3bf-f5be89229e70.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/225324db-5d51-4f69-ab95-ccbcd89717b7.png.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "kakbah-blue-sky-case-hp-kaligrafi-nama-tersedia-semua-tipe-hp-ultimatte-black-1b768",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 1",
    "country": "Kab. Cianjur",
    "estimation": "Estimasi tiba 14 - 17 Nov",
    "toko": {
      "username": "ahvisunday",
      "toko_name": "AhviSunday",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2022/6/7/a74cd2c3-c2fd-40fa-a651-72321b1b6bf4.jpg"
    }
  },
  {
    "product_name": "null",
    "price": "Rp134.850",
    "stock": "19",
    "description": "Official Store Ringke<br>MOHON PERHATIKAN VARIAN PRODUK SEBELUM MELAKUKAN PEMBELIAN<br>TIDAK MENERIMA ORDER VIA CHAT ATAU NOTES<br>SALAH PEMESANAN TIDAK BISA RETUR <br><br>Varian HP:<br>- iPhone 12 / 12 Pro<br>- iPhone 12 Pro Max<br>- iPhone 12 Mini<br><br>Fitur Produk:<br>* Mag Safe Compatibel<br><br>* Melindungi Kamera dan Layar (Front Lip Raise)<br><br>* Material: Softcase TPU<br><br>* Slim, Fashionable dan Minimalis<br><br>* Melindungi dari debu, kotor, goresan<br><br>* Simpel dan Presisi<br><br>* Tombol Mudah di Akses<br><br>* Nyaman di genggam<br>Design minimalis dengan warna transparan sehingga tetap menonjolkan design asli perangkat. Case Slim dan Ringan, tidak membuat perangkat menjadi bulky, cocok untuk yang menyukai case dengan design simpel<br><br>* Raised Bezel <br>Frame depan casing dibuat sedikit timbul untuk melindungi layar dari gesekan<br><br>* Kompatibel Wireless Charging<br>Tidak perlu melepas Case saat mengisi daya dengan Qi Wireless charging<br><br>* Potongan Akurat dan Tombol Presisi<br><br>* Bisa dipakai dengan beragam jenis screen protector dan Ringke Dual Easy Film<br><br>Why Buy From Us:<br>*Official Distributor<br>*Official Store<br>*Garansi Original 2x Money Back Guarantee<br>*Garansi Tukar Barang Baru Jika Ada Cacat Pabrik Selama 2 Hari dari penerimaan<br>*Customer Satisfaction Guaranted",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/5/28/d80a8dd2-98fd-46d9-96b9-b5b268b8453f.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/5/28/c6f78ba0-ebf7-4a1f-a788-8546424a2d90.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/5/28/c074161f-0c1a-4392-811d-b2807d2a3913.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/5/28/5acab18b-0fd3-4123-a29a-40feb2dc3cc4.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/5/28/cf5032ce-eb80-49a4-8ebb-2c20a581f7d7.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/5/26/70778d6b-023a-4d25-a1b3-705f3dcb9473.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/5/26/27a90a6f-f5d1-4975-81f6-cf96b067af83.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/5/26/cf8a78d5-2aa1-46d6-88e8-dfbaab75f217.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/5/26/1389c9e1-892c-43a0-8563-afd600dd89b1.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/5/26/7122e8e2-deae-4c3d-838d-e968f711b24a.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/5/26/5764c0de-c223-4ee1-90dc-f8d3ec713a6c.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/4bff5c36-8741-429b-94f2-90d959672265.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/62044c2f-f0d3-43a5-ac5a-41cf6c6f8e92.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/ecc85091-66d0-4075-8350-3ad6c54f1d12.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/3595c2e3-e9d1-45d5-9ad3-74b0320ecd9d.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/7da85b26-f5e0-4d5d-9f6e-acb641e2d7a3.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/4bdc3190-264f-48f9-a231-e76b47fbde3a.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/0dd3df03-4fb0-4cd8-bfcf-25c1aed6a582.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/5b5e3c70-31f3-4ab5-a16e-708387aad1b0.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/ec2991c4-4484-42bb-bd53-55bd447ce773.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/5021619c-f0ec-43d1-98e3-2e437f3da33f.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/4bdc3190-264f-48f9-a231-e76b47fbde3a.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/0dd3df03-4fb0-4cd8-bfcf-25c1aed6a582.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/5b5e3c70-31f3-4ab5-a16e-708387aad1b0.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/ec2991c4-4484-42bb-bd53-55bd447ce773.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/5021619c-f0ec-43d1-98e3-2e437f3da33f.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/24/c9d12db4-d6b5-4f8c-a9a5-8f00b64a6b15.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/24/8c25d5ed-9e7e-41f4-a581-a9ba205265a0.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/24/f03f9846-08cd-4c1b-8eef-f163aa7bbdf6.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/24/10fac142-3d0a-4cc3-9cc4-8896993e3386.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/24/012ab88b-9ed1-44e6-b53b-65b0f520f513.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/24/a940486b-42f3-4e0c-a325-408d43c6d025.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/24/c9d12db4-d6b5-4f8c-a9a5-8f00b64a6b15.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/24/8c25d5ed-9e7e-41f4-a581-a9ba205265a0.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/24/f03f9846-08cd-4c1b-8eef-f163aa7bbdf6.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/24/10fac142-3d0a-4cc3-9cc4-8896993e3386.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/24/012ab88b-9ed1-44e6-b53b-65b0f520f513.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/24/a940486b-42f3-4e0c-a325-408d43c6d025.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/22/44577a4f-afca-4d25-b3e5-faae653d24f4.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/22/21394409-0daa-45e3-a7cd-b44d297738f6.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/22/7ecf734c-53a7-4c79-b5ad-aa8ee05a09d5.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/22/5096a4cc-6d4f-46a0-b29d-db868e0f2df5.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/22/f36927a1-9c6e-4513-aeaa-129e4cf0c959.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/22/d463ae6f-b70e-44fa-bc13-653e5eae0f03.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/14/95c0a70e-af6f-49af-bc21-8bc6020e1c19.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/14/b9d0e373-1763-46c5-b5d4-257c989cd2e2.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/14/d5419893-6b75-41ae-9913-b2296fd7b492.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/14/ff73b626-7490-4579-bf84-e9a8ae34e0f1.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/14/81f6efe5-8f0f-4fa5-8e1c-ee7c7ef6348d.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/14/f3146b76-73b5-4583-ae0d-a64a92de9356.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "case-ringke-iphone-12-mini-12-12-pro-12-pro-max-air-s-black-pink-gray-casing-softcase-anti-crack-tipis-slim-armor-military-drop-tahan-banting-pelindung-handphone-sarung-hp-tpu-cover-silikon-hardcase-clear-bening-jernih-transparan-12m-black-2bab0",
    "discount": "51%",
    "discount_price": "Rp273.000",
    "rating": "5",
    "sold": "Terjual 2",
    "country": "Jakarta Utara",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "officialstoreringke",
      "toko_name": "Official Store Ringke",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/5/22/db64de0a-778b-4ff5-86ea-30ad69cd3eff.jpg"
    }
  },
  {
    "product_name": "Casing Hp iPhone X Xr Xs Xs Max Soft Case Liquid Magnetic Ring Holder - Green, Iphone xs",
    "price": "Rp45.000",
    "stock": "994",
    "description": "Soft Case Silicone Liquid and Ring Holder (Free Tali / Sesuai Gambar Produk)<br><br>Ready Warna:<br>* Black / Hitam<br>* Blue / Biru<br>* Purple / Ungu<br>* Green / Hijau<br><br>NB:<br>Mohon cantumkan warna cadangan di keterangan produk, bila warna yang kakak pilih tidak tersedia<br><br>Deskripsi Produk :<br>- Ring holder bisa diputar, bisa dibuat untuk stand hp<br>- Casing lebih tinggi 0.3mm dari kamera belakang dan 0.3mm dari layar depan<br>- Material liquid silicone<br>- Mudah dibersihkan dari debu<br>- Bagian dalam halus, body belakang hp dijamin aman tidak tergores<br>- Mudah dipasang dan dilepas<br>- Dijamin presisi sesuai type handphone<br>- Case friendly. Aman dipasang dengan segala jenis anti gores.<br><br>Terima kasih,",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/12/27/694e7e99-2472-4560-84af-398cab716cd1.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/12/27/d0b730cd-07b6-49d6-987b-f41ec6b6c63a.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/12/27/10582270-d171-4513-bdc7-ecbfc9278182.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/12/27/198ad21a-e41e-4e97-a8bf-dcef773ff495.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-iphone-x-xr-xs-xs-max-soft-case-liquid-magnetic-ring-holder-green-iphone-xs",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 90+",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "fahryonemore",
      "toko_name": "FahryOneMore",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/9/7/3bb0f3ba-2d30-424a-885a-b4599e243647.jpg"
    }
  },
  {
    "product_name": "Casing Hp Tecno Pova 5 Pro Soft Case Matte Anti Licin Premium - Biru",
    "price": "Rp28.000",
    "stock": "83",
    "description": "NEW ARRIVAL<br>Premium Case Anti Fingerprint<br><br>Matte Case <br><br>Ready Colour:<br>* Black<br>* Blue Navy<br><br>Type: Carbon Case<br>Function: Dirt-resistant<br>Compatible Brand: Tecno<br>Design: Plain,Matte,Business,Vintage<br>Features: Real Matte Anti Fingerprint TPU Case<br><br>Spesifikasi &amp; Review Singkat :<br>* Produk Signature<br>* Desain Ultra tipis 0.6 mm dan perfect fit<br>* Desain minimalis dan maskulin<br>* Detail sangat tinggi dan sangat presisi<br>* Sangat ringan, lentur, tidak mudah kendor<br>* Sangat nyaman digenggam dalam waktu lama<br>* Lubang dan sisi tertutupi sangat rapi<br>* Melindungi benturan dari semua sisi<br>* Ciri khas dengan lis UV Edge Finish<br>* Warna tajam yang tahan lama (anti kusam)<br>* Material Solid TPU berkualitas terbaik<br>* Matte Finish yang halus sempurna<br>* Sisi ujung lebih keras untuk perlindungan lebih<br>* Tidak menganggu sinyal ponsel<br>* tidak membuat baret pada ponsel<br>* Coating Anti Jamur, Bebas jamur<br>* Coating Anti Minyak, Bebas bekas jari<br>* Coating Anti Kotor, Mudah Dibersihkan<br>* Coating Anti Gores (Sewajarnya)<br><br>Terima kasih.",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/6/fc197e3d-004c-4481-990e-9d61b937779b.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/6/065d417f-3eb6-4745-a20e-c72053e1779b.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/6/a699ee84-d1a1-4271-96e9-0f171dac18d6.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/6/1058f692-aeb4-42b6-b269-0b4236c7a83d.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-tecno-pova-5-pro-soft-case-matte-anti-licin-premium-biru-f7c70",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.9",
    "sold": "Terjual 100+",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "fahryonemore",
      "toko_name": "FahryOneMore",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/9/7/3bb0f3ba-2d30-424a-885a-b4599e243647.jpg"
    }
  },
  {
    "product_name": "Casing Samsung Galaxy Z Fold 6 5 4 5G Silicone Case Cover SoftCase Pelindung Hp Z Fold6 Fold5 - Fold6 White",
    "price": "Rp749.000",
    "stock": "3",
    "description": "BARANG ORIGINAL !!!!<br><br>CHECK JUGA GADGET FAVORITMU DISINI<br>Smarphone : <a target=\"_blank\" rel=\"noopener noreferrer\" href=\"https://www.tokopedia.com/supergadgettt/etalase/smartphone\">https://www.tokopedia.com/supergadgettt/etalase/smartphone</a><br>Tablet : <a target=\"_blank\" rel=\"noopener noreferrer\" href=\"https://www.tokopedia.com/supergadgettt/etalase/tablets\">https://www.tokopedia.com/supergadgettt/etalase/tablets</a><br><br>Dapat digunakan untuk :<br>Samsung Galaxy Z Fold4 | Fold5 | Fold6<br><br>Silicone Case<br><br>Spesifikasi:<br>1. Tampilan desain premium<br>2. Proteksi tertinggi untuk ponsel<br>3. Dibuat dengan bahan yang ramah lingkungan<br><br><br>Notes<br>- Pastikan sudah memilih varian model sudah tepat dan sesuai dengan model hp anda, karena judul produk dibuat secara universal untuk semua tipe hp.<br><br>** Perhatian Varian Model / Warna Produk Yang Dipilih Sebelum Melakukan Pembelian",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/13/9ae93211-ff40-4785-874c-4179c4d78c2d.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/13/1d862cf2-49ce-4f22-aec4-b286c90a31b0.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/13/c53131f9-1268-40aa-9abf-44f9fb74a7ed.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/13/19e7738b-586b-44ac-9133-cc6e280a2f83.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/8/16/89376435-8a1e-4a01-ad48-c3b00df1f1c2.png.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-samsung-galaxy-z-fold-6-5-4-5g-silicone-case-cover-softcase-pelindung-hp-z-fold6-fold5-fold6-white-e418a",
    "discount": "6%",
    "discount_price": "Rp799.000",
    "rating": "4.9",
    "sold": "Terjual 27",
    "country": "Jakarta Pusat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "supergadgettt",
      "toko_name": "SUPER_GADGET",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2019/8/8/405517/405517_a23f3177-3e6a-44cb-a01f-ff132e8cea0c.png"
    }
  },
  {
    "product_name": "NEW FANCY MIRROR case iPhone 13 14 15 Pro Max Softcase Casing HP Cover - Lovely, iPh 13",
    "price": "Rp37.500",
    "stock": "20",
    "description": "⚠️ PENTING ⚠️<br>‼️ MOHON DIPERHATIKAN ‼️<br>Tolong pilih VARIASI PRODUK sesuai type HP yang dibutuhkan<br>Pesanan akan dikirim sesuai invoice bukan sesuai chat atau diskusi<br>Barang yang sudah dikirim sesuai pesanan tidak dapat ditukar<br><br>DESKRIPSI PRODUK :<br>✅ Bahan soft TPU Jelly Transparant Lentur<br>✅ Ketebalan +-3mm Sangat melindungi<br>✅ Dengan 3 pilihan warna yang sangat menarik<br>✅ Camera Protect<br>✅ Body tercover sempurna<br>✅ Tombol Empuk dan Lubang-Lubang Presisi<br>✅ Dengan kaca asli pada bagian belakang case",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/2/28/a9eb9741-e209-4694-b749-c560f4229c14.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/2/28/922b96df-01cb-44e6-86bf-68d0ce8398ad.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/2/28/84e53dfa-e9f0-411d-8bea-8264c2b6bcc0.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/2/28/154c2cd3-84e1-4ce9-901b-f844fd12d6f6.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/2/28/62aa347a-088d-4e0b-81d3-a367ea8f4ca4.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "new-fancy-mirror-case-iphone-13-14-15-pro-max-softcase-casing-hp-cover-lovely-iph-13-46658",
    "discount": "25%",
    "discount_price": "Rp50.000",
    "rating": "5",
    "sold": "Terjual 50+",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "casepedias",
      "toko_name": "Case Pedia",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2017/1/5/1042576/1042576_4d8ee5d1-ccf0-4adc-aac9-5fbf9eeadc99.jpg"
    }
  },
  {
    "product_name": "Casing HP Samsung A52 A50 A51 A71 A70 A72 A7 2018 Kartun Bunga-bunga - Bright Green, Samsung A50",
    "price": "Rp25.000",
    "stock": "990",
    "description": "Product Description:<br>100% Brand New and Removeable.<br>Fashion square pattern design keeps your phone private and feels good.<br>Shock absorbent, shatterproof material.<br>Lanyard added protection against accidental drops.<br>All function works well even your phone was in this case.<br><br>Feature:<br>Material: Made of High Quality Silicone + TPU<br><br>Package Content:<br>1 x Plating Case + 1 x Lanyard",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/25/8d63aed6-bf6b-4353-b6d4-15e17df71730.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/25/32e32991-a86b-4b42-aac8-caad931135f0.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/25/bff54914-50d6-4945-9c86-ea86d6749873.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/25/f65ef3df-6c30-408e-b78f-dea3feca84aa.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/25/824d4cbe-c9ba-404b-83c7-004a1e2c95b6.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-samsung-a52-a50-a51-a71-a70-a72-a7-2018-kartun-bunga-bunga-bright-green-samsung-a50",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.9",
    "sold": "Terjual 100+",
    "country": "Kab. Tangerang",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "shopeper",
      "toko_name": "Trendi Case",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/4/29/1df532de-5202-4b3e-8db7-a42d37a779de.jpg"
    }
  },
  {
    "product_name": "Casing Handphone Glitter Ribbon Clear Case 11 12 13 14 15 PRO MAX - 13 PROMAX",
    "price": "Rp68.000",
    "stock": "91",
    "description": "Tersedia Mulai dari : XR,11 s.d 15 Series<br><br>Fitur lainnya : <br>- Bahan TPU clear berkualitas tinggi<br>- Design ribbon pada bagian tengah body membuat case terlihat mewah<br>- Selain itu, terdapat aksen glitter yang mengkilap<br>- Mudah dipasang &amp; dikeluarkan<br>-\\u2060 Memberikan proteksi pada bagian kamera, karena lips bagian kamera sudah lebih tinggi daripada lensa<br>- Tipe XR dan 11 untuk case ini satu tipe ya Kak. Pakai tipe 11<br><br>Untuk case ribbon ini case dengan tema dan kelebihan product DIY (handmade) sehingga yg lebih ditonjolkan pada product bukan pada bagian presisi dan atau kerapihan penempatannya. tetapi memang lebih ke handmadenya di tempel satu persatu dan dikeringkan secara manual sehingga product ini tidak akan sama antara satu dengan yg lainnya. Berbeda dengan diprint oleh mesin yg pastinya akan presisi dan sama😊<br><br>Notes : bagian pita tidak disarankan terkena alcohol atau benda kasar karna akan berubah warna<br><br>Pilihan Varian :<br>- Clear Ribbon<br><br>**KETERANGAN<br><br>- Silahkan diorder , Selama tipe HP dan warna masih bisa dipilih, artinya stock barang kami masih READY. Mohon kepada customer untuk memilih tipe HP dan warna yang sesuai untuk mempercepat proses penyiapan barang dan pengiriman<br>- Jika tipe HP dan varian warna tidak bisa dipilih, artinya stock kami sedang kosong . Jika butuh bantuan/pertanyaan silahkan chat MINZEN ya<br>Happy Shopping and Have a Nice Day ^^",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/2/23/912f7a60-8e7d-4bce-a057-7c5ec708ae8d.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/2/23/31e41a9f-b78f-4412-848c-9776f684db8d.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/11/dad25893-7fbd-472c-9f04-ab39b79e8565.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/11/b99e0b15-acdb-41b9-95d9-a0d3a5e26cf9.png.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-handphone-glitter-ribbon-clear-case-11-12-13-14-15-pro-max-13-promax-d5bad",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.8",
    "sold": "Terjual 50+",
    "country": "Kota Tangerang",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "ntzenonline",
      "toko_name": "NTzen Online",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2024/9/27/f0687b61-eebb-48c5-8887-204df07d7986.png"
    }
  },
  {
    "product_name": "Case iPhone Ultrathin Tipis TPU Slim casing hp - IP 7/8/SE, Note Warna/Ctt",
    "price": "Rp25.000",
    "stock": "10",
    "description": "-- ANTI KUNING CASE ! CASING INI TIDAK AKAN BISA MENGUNING !!--<br><br>Casing yang termasuk softcase DAN transparan ini hanya memiliki ketebalan 0.2 MM Saja !! Cocok bagi kalian yang suka casing simple , keren , cool , dan elegan .<br><br><br>ANDA MAU MERASAKAN MEMEGANG HP SEPERTI TANPA CASING ?RASAKAN FEEL KETIKA MEMAKAI HP SEPERTI TANPA PAKAI CASING !! Casing ini anti minyak / Fingerprint tentunya ! Sehingga HP kalian akan tetap nyaman untuk digenggam .<br><br>Casing ini juga memiliki design yang sangat unik serta simple . RASAKAN grip yang sangat nyaman untuk HP anda ketika memakai casing ini<br><br>- PURPLE ( WARNA BARU )<br>- Black<br>- Grey<br>- Blue<br>- Midnight Green<br>- White<br>TIPE HP :<br>- 6 / 6S<br>- 6 PLUS / 6S PLUS<br>- 7 / 8 / SE 2020<br>- 7 PLUS / 8 PLUS<br>- X / XS<br>- XR<br>- XS MAX<br>- 11<br>- 11 PRO<br>- 11 PRO MAX<br>- 12<br>- 12 PRO<br>- 12 PRO MAX<br>- 12 MINI<br>- iPHONE 13/13PRO/13PROMAX<br>- IPHONE 14/14PRO/14PROMAX/14PLUS<br>- IPHONE 15/15PRO/15PROMAX/15PLUS<br><br>-----------------------------------<br>*ORDER SEBELUM JAM 21.00 AKAN DIKIRIM DIHARI YANG SAMA<br>*GOJEK &amp; GRAB INSTANT/SAMEDAY READY SELAMAT BERBELANJA:)<br>",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/2/7bf14e34-22c9-42f2-9f0e-04505c8294a8.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/2/b987bd0d-729a-406f-8feb-e8399b766c20.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/2/07f7fb43-22ac-48a6-8c0b-6c7488923c13.png.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "case-iphone-ultrathin-tipis-tpu-slim-casing-hp-ip-7-8-se-note-warna-ctt-3ac8b",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 14",
    "country": "Kota Surakarta",
    "estimation": "Estimasi tiba 31 Okt - 2 Nov",
    "toko": {
      "username": "kesing99",
      "toko_name": "Kesing99_online",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/3/18/8bbfb747-cca8-4e44-9f2e-d7735d332dec.jpg"
    }
  },
  {
    "product_name": "Soft Case Xiaomi Mi A1 Kesing Casing Silikon Silicon Cover Kondom HP - Hijau",
    "price": "Rp7.900",
    "stock": "11",
    "description": "💛100% Gambar Asli / Real Pictures💛<br>💜100% Gambar Asli / Real Pictures💜<br><br>- Terbuat dari bahan Silicon Premium<br>- Terdapat pelindung dibagian kamera<br>- Tidak mudah kotor dan berubah warna<br>- Permukaan lembut dan tidak licin<br>- Tahan benturan, Kuat dan Awet<br>- Elastis dan lentur tapi tidak mudah melar!!<br><br> &gt;&gt;Tersedia 7 warna :<br>- Hitam<br>- Biru<br>- Coklat<br>- Hijau<br>- Rose (Pink Kecoklatan)<br>- Maroon<br>- Broken White<br><br>**sertakan pilihan warna cadangan pada catatan<br>**jika warna yg dipilih kosong akan kami konfirmasi",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/11/29/3ad7bf97-032b-43dc-8e85-f4131aaaa5e7.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/11/29/7ef31bb8-f1d7-4f56-843c-135dda5f1c42.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/11/29/0b1df815-5c59-4d86-a4d2-dc32ceb487a3.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/11/29/5ad20f87-6c76-469b-b666-a9bf39ff6d2d.png.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "soft-case-xiaomi-mi-a1-kesing-casing-silikon-silicon-cover-kondom-hp-hijau",
    "discount": "68%",
    "discount_price": "Rp24.900",
    "rating": "5",
    "sold": "Terjual 100+",
    "country": "Kota Bekasi",
    "estimation": "Estimasi tiba besok - 2 Nov",
    "toko": {
      "username": "rajaapasaja",
      "toko_name": "Raja Apa Saja",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/4/29/41f8fb71-cdd2-475f-9cfd-beea726032dc.png"
    }
  },
  {
    "product_name": "Casing HP iPhone 13 Aesthetic Premium Leather Aegis Fawn Woody Brown",
    "price": "Rp1.050.000",
    "stock": "3",
    "description": "Aegis Fawn made of genuine leather. You'll find the textured grain feels so soft and will form a natural patina colour over time. Our cases are fit designed and easy-to-press buttons. Also, the raised-edge will keep your phone protected. This one-of-a-kind accessory will never go out of style.<br><br>Specification:<br>- Genuine Leather<br>- Compatibility: iPhone 13<br>- Develops patina over time<br>- Gold debossed<br><br><br><br>Aegis Fawn terbuat dari kulit asli. Anda akan menemukan butiran bertekstur terasa begitu lembut dan akan membentuk warna patina alami seiring waktu. Kasing kami dirancang dengan pas dan tombol yang mudah ditekan. Selain itu, tepi yang ditinggikan akan menjaga ponsel Anda tetap terlindungi. Aksesori unik ini tidak akan pernah ketinggalan zaman.<br><br>Spesifikasi:<br>- Kulit Asli<br>- Kompatibilitas: iPhone 13<br>- Mengembangkan patina dari waktu ke waktu<br>- Debossed emas",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/5/5b8d102a-372c-4a63-b25a-696da6cf87f1.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/5/0fc4d186-d89d-4a8b-8c86-ae46431f8d4b.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/5/c8151aa6-3b86-42ad-8e87-b1fb19a6d8c3.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/5/cb942289-2b44-4890-8aef-89e83f7d355e.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/5/7c203f53-8013-43ad-8198-00933c3a8e71.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/5/5b8d102a-372c-4a63-b25a-696da6cf87f1.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/5/5b8d102a-372c-4a63-b25a-696da6cf87f1.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/5/0fc4d186-d89d-4a8b-8c86-ae46431f8d4b.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/5/c8151aa6-3b86-42ad-8e87-b1fb19a6d8c3.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/5/cb942289-2b44-4890-8aef-89e83f7d355e.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/5/7c203f53-8013-43ad-8198-00933c3a8e71.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-iphone-13-aesthetic-premium-leather-aegis-fawn-woody-brown",
    "discount": "null",
    "discount_price": "null",
    "rating": "null",
    "sold": "null",
    "country": "Jakarta Utara",
    "estimation": "Estimasi tiba besok - 2 Nov",
    "toko": {
      "username": "starke-leather",
      "toko_name": "Starke Leather Official Store",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2016/3/16/8031674/8031674_f41367e5-00d4-476a-bb6e-f8d81402b3c7.jpg"
    }
  },
  {
    "product_name": "Casing HP Kulit Asli Phone 13 Premium Leather Aegis Fawn Tawny Brown",
    "price": "Rp1.050.000",
    "stock": "3",
    "description": "Aegis Fawn made of genuine leather. You'll find the textured grain feels so soft and will form a natural patina colour over time. Our cases are fit designed and easy-to-press buttons. Also, the raised-edge will keep your phone protected. This one-of-a-kind accessory will never go out of style.<br><br>Specification:<br>- Genuine Leather<br>- Compatibility: iPhone 13<br>- Develops patina over time<br>- Gold debossed<br><br><br><br>Aegis Fawn terbuat dari kulit asli. Anda akan menemukan butiran bertekstur terasa begitu lembut dan akan membentuk warna patina alami seiring waktu. Kasing kami dirancang dengan pas dan tombol yang mudah ditekan. Selain itu, tepi yang ditinggikan akan menjaga ponsel Anda tetap terlindungi. Aksesori unik ini tidak akan pernah ketinggalan zaman.<br><br>Spesifikasi:<br>- Kulit Asli<br>- Kompatibilitas: iPhone 13<br>- Mengembangkan patina dari waktu ke waktu<br>- Debossed emas",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/5/5a94211c-29a9-47b0-83aa-3366534a80cc.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/5/f48282a0-91ca-42ae-8327-ff46a247cb6a.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/5/2db91bde-7d62-43de-8c8a-e619b0d39004.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/5/45e0e10c-d019-4ece-a758-59ae45a1a696.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/10/5/e7eb7bba-6ed8-41c5-8d41-195acc16bc48.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-kulit-asli-phone-13-premium-leather-aegis-fawn-tawny-brown",
    "discount": "null",
    "discount_price": "null",
    "rating": "null",
    "sold": "null",
    "country": "Jakarta Utara",
    "estimation": "Estimasi tiba besok - 2 Nov",
    "toko": {
      "username": "starke-leather",
      "toko_name": "Starke Leather Official Store",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2016/3/16/8031674/8031674_f41367e5-00d4-476a-bb6e-f8d81402b3c7.jpg"
    }
  },
  {
    "product_name": "Kakbah (Afternoon Sky) - Case Hp Kaligrafi Nama (Tersedia Semua Tipe HP) - Ultimatte Black",
    "price": "Rp109.000",
    "stock": "99",
    "description": "AhviSunday&nbsp;adalah&nbsp;brand&nbsp;Custom&nbsp;Kaligrafi&nbsp;terlaris&nbsp;Di&nbsp;Indonesia.&nbsp;Ribuan&nbsp;pesanan&nbsp;untuk&nbsp;case&nbsp;HP&nbsp;sudah&nbsp;dipercayakan&nbsp;ke&nbsp;AhviSunday.<br><br>🔥&nbsp;Tersedia&nbsp;semua&nbsp;tipe&nbsp;HP&nbsp;:&nbsp;Vivo,&nbsp;Oppo,&nbsp;Xiaomi,&nbsp;Lenovo,&nbsp;Iphone,&nbsp;Samsung&nbsp;dll<br>🔥&nbsp;Kualitas&nbsp;Mesin&nbsp;UV&nbsp;asli&nbsp;Jepang<br>🔥&nbsp;Hasil&nbsp;cetak&nbsp;jelas&nbsp;dan&nbsp;tidak&nbsp;luntur<br>🔥&nbsp;Bahan&nbsp;ringan&nbsp;dalam&nbsp;genggaman<br>🔥&nbsp;Finishing&nbsp;Doff&nbsp;/&nbsp;Glossy<br><br>==================================================<br><br>Informasi&nbsp;Bahan&nbsp;:<br><br>Ultimatte&nbsp;Black<br>✔️&nbsp;Bahan&nbsp;dasar&nbsp;silikon<br>✔️&nbsp;Cetak&nbsp;bagian&nbsp;permukaan&nbsp;belakang<br>✔️&nbsp;Lentur&nbsp;dan&nbsp;fit&nbsp;dengan&nbsp;smartphone<br>✔️&nbsp;Dasar&nbsp;Case&nbsp;berwarna&nbsp;hitam<br><br>Clear&nbsp;Case<br>✔️&nbsp;Bahan&nbsp;dasar&nbsp;silikon&nbsp;bening<br>✔️&nbsp;Cetak&nbsp;bagian&nbsp;permukaan&nbsp;belakang<br>✔️&nbsp;Lentur&nbsp;dan&nbsp;fit&nbsp;dengan&nbsp;smartphone<br>✔️&nbsp;Elastis&nbsp;bisa&nbsp;kembali&nbsp;ke&nbsp;bentuk&nbsp;semula<br>✔️&nbsp;Ada&nbsp;sisi&nbsp;bumper&nbsp;di&nbsp;sudut&nbsp;untuk&nbsp;meminimalisir&nbsp;retak&nbsp;apabila&nbsp;terjatuh<br><br>Slim&nbsp;Clear&nbsp;Case<br>✔️&nbsp;Bahan&nbsp;dasar&nbsp;silikon&nbsp;bening<br>✔️&nbsp;Cetak&nbsp;bagian&nbsp;permukaan&nbsp;belakang<br>✔️&nbsp;Lentur&nbsp;dan&nbsp;fit&nbsp;dengan&nbsp;smartphone<br>✔️&nbsp;Elastis&nbsp;bisa&nbsp;kembali&nbsp;ke&nbsp;bentuk&nbsp;semula<br><br>Exclusive&nbsp;Case<br>✔️&nbsp;Anti&nbsp;Gores,&nbsp;Anti&nbsp;Api&nbsp;dan&nbsp;Anti&nbsp;Noda<br>✔️&nbsp;Full&nbsp;Glass&nbsp;/&nbsp;plat&nbsp;besi&nbsp;dengan&nbsp;Coating&nbsp;TPU<br>✔️&nbsp;Print&nbsp;pada&nbsp;permukaan<br>✔️&nbsp;Menggunakan&nbsp;sistem&nbsp;cetak&nbsp;Transfer&nbsp;Printing&nbsp;Film<br>✔️&nbsp;Bevel&nbsp;di&nbsp;ujung&nbsp;case&nbsp;agar&nbsp;glass&nbsp;tidak&nbsp;tajam<br><br>Fullprint&nbsp;Hardcase<br>✔️&nbsp;Bahan&nbsp;Polycarbonate<br>✔️&nbsp;Full&nbsp;print&nbsp;/&nbsp;cetak&nbsp;hingga&nbsp;sisi&nbsp;samping&nbsp;case<br>✔️&nbsp;Hasil&nbsp;jelas&nbsp;dan&nbsp;tidak&nbsp;luntur<br>✔️&nbsp;Atas&nbsp;dan&nbsp;bawah&nbsp;case&nbsp;terbuka<br>✔️&nbsp;Bahan&nbsp;ringan&nbsp;di&nbsp;dalam&nbsp;genggaman&nbsp;walau&nbsp;bahan&nbsp;keras<br><br>==================================================<br><br>Cara&nbsp;Pemesanan&nbsp;:<br><br>⚠️&nbsp;PENTING!!!&nbsp;Tanyakan&nbsp;ketersediaan&nbsp;stok&nbsp;case&nbsp;terlebih&nbsp;dahulu<br>✔️&nbsp;Pilih&nbsp;jenis&nbsp;case<br>✔️&nbsp;Masukkan&nbsp;gambar&nbsp;desain&nbsp;pada&nbsp;keterangan&nbsp;/&nbsp;chat&nbsp;via&nbsp;tokopedia<br>✔️&nbsp;Mohon&nbsp;cantumkan&nbsp;nama&nbsp;kaligrafi&nbsp;beserta&nbsp;tipe&nbsp;hpnya&nbsp;dicatatan&nbsp;pembelian<br><br>FREE&nbsp;Garansi&nbsp;/&nbsp;Retur&nbsp;Jika&nbsp;diterima&nbsp;kondisi&nbsp;rusak",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/c27654e0-e5a5-4098-80b0-d5790686f2cd.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/dd847b5f-9076-4fd7-89cb-2c387751f4b8.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/70332ce2-ba6c-4af6-b9f4-032eb12187eb.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/c27654e0-e5a5-4098-80b0-d5790686f2cd.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/03023331-dd6e-4f7c-bfcf-70ac05605e8e.png.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "kakbah-afternoon-sky-case-hp-kaligrafi-nama-tersedia-semua-tipe-hp-ultimatte-black-be223",
    "discount": "null",
    "discount_price": "null",
    "rating": "null",
    "sold": "1 barang berhasil terjual",
    "country": "Kab. Cianjur",
    "estimation": "Estimasi tiba 14 - 17 Nov",
    "toko": {
      "username": "ahvisunday",
      "toko_name": "AhviSunday",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2022/6/7/a74cd2c3-c2fd-40fa-a651-72321b1b6bf4.jpg"
    }
  },
  {
    "product_name": "CASE SAMSUNG A53 A33 A73 5G SILIKON BLACK RING COVER SOFT CASE HP - HITAM + RING, SAMSUNG A53",
    "price": "Rp27.000",
    "stock": "92",
    "description": "CASE LEATHER PU&nbsp;<br>SMART CASE&nbsp;<br><br><br>BAHAN ; TPU SILIKON MEWAH DENGAN LAPISAN KULIT YG SANGAT RAPIH DAN ELEGANT&nbsp;<br><br>WARNA ; HITAM <br><br>SEMUA BARANG SUDAH MELALUI QC DAN JAMIN PAS DI HP , SEMUA FOTO YG KAMI TAYANG KAN ADALAH FOTO REAL DAN SESUAI WARNA .&nbsp;<br><br>APABILA TERJADI KERUSAKAN BUKAN TANGGUNG JAWAB KAMI DAN SEMUA KMPLAIN HARUS ADA UNBOXING VIDEO&nbsp;<br><br>UNTUK KEAMANAN BARANG DI ANJURKAN UNTUK MEMBELI BOX DAN BUBBLE&nbsp;<br><br>THANKS&nbsp;",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/3/15/95df62ef-ce03-48a2-a81a-bda85f9207a1.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/4/19/52a55bcd-9740-4f52-b9c1-55410de5161e.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/4/13/35b9096d-358f-4e82-bf78-d020d3cb0b3c.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/3/24/0a9f0851-1117-4757-aa20-9792c6812dc0.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/3/15/83b6e97b-e0ec-440c-befa-29cdf4b4ae49.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "case-samsung-a53-a33-a73-5g-silikon-black-ring-cover-soft-case-hp-hitam-ring-samsung-a53",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 100+",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "pamola",
      "toko_name": "PAMOLLA",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/3/6/a2c043a9-b0e5-4bbe-b1f5-86e7096e2443.jpg"
    }
  },
  {
    "product_name": "Case hp iphone 12 / 12 pro / 12 pro max tpu metal magsafe premium - Biru, Ip 12pro",
    "price": "Rp85.000",
    "stock": "100",
    "description": "tersedia type <br>Iphone 11<br>Iphone 11 pro<br>Iphone 11 pro max<br>Iphone 12 <br>Iphone 12 pro<br>Iphone 13<br>Iphone 13 pro <br>Iphone 13 pro max<br>Iphone 14 <br>Iphone 14 plus<br>Iphone 14 pro <br>Iphone 14 pro max<br><br>Tersedia warna<br>Hijau<br>Bening<br>Biru<br>Hitam<br>Ungu<br><br>Matte Magsafe magnetic wireless<br>Charging phone Case menampilkan tekstur matte halus yang memberikan tampilan canggih, dan cengkeraman yang ditingkatkan untuk mencegah terjatuh secara tidak sengaja.<br><br>Dirancang dengan teknologi MagSafe, Matte Hybrid Case terhubung dengan mulus ke pengisi daya dan aksesori MagSafe dan dengan penutup tombol aluminium taktil, memastikan fungsionalitas tanpa batas, perlindungan tertinggi, dan kemudahan penggunaan.<br><br>• Premium Matte Finish<br>• High Strength Polycarbonate back panel<br>• 2H scratch resistance<br>• Rugged Stealth Design<br>• Tactile responsive Aluminium buttons<br>• Superb Fit, Design, Feel &amp; Quality<br>• Raised camera lip protects the camera against scratches<br>• Full Coverage Protection (Sides and Buttons)<br>• Strong Magsafe Magnets<br><br>Aman untk case iphone<br>Tpu metal magsafe magnetic case shockproof solid bumper cover back matte anti sidik jari",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/24/e6552958-ae63-433e-8a26-39769fb1df93.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/24/86230e7e-cabd-4a9a-bbb1-714cd18d0f45.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/24/912bab14-32a2-4eff-9441-c19ce6955ed4.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/24/22e46e65-7030-44f1-ac25-174b6838c92c.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/24/e6552958-ae63-433e-8a26-39769fb1df93.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/24/c38ce507-fd8b-40c3-8615-9d482f509213.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/24/86230e7e-cabd-4a9a-bbb1-714cd18d0f45.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/24/912bab14-32a2-4eff-9441-c19ce6955ed4.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/24/22e46e65-7030-44f1-ac25-174b6838c92c.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "case-hp-iphone-12-12-pro-12-pro-max-tpu-metal-magsafe-premium-biru-ip-12pro-219b7",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 5",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "sanelvisacc",
      "toko_name": "Sanelvisacc",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2022/6/23/8422baec-4e21-4bd7-8421-ba75ade44750.jpg"
    }
  },
  {
    "product_name": "Case hp iphone 13 / 13 pro / 13 pro max softcase magsafe premium - Black, Iphone 13pro",
    "price": "Rp49.500",
    "stock": "100",
    "description": "wellcome<br><br>soft case hp iphone magsafe mewah<br><br>tersedia warna<br><br>- Black<br>- Pink<br>- Navy<br>- purple<br><br>happy shopping",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/1/21/e624fdc4-c4f2-4f63-b9e7-19f45d894afb.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/1/21/ae8430ad-1afb-4d0d-9e8c-c1c978ed7504.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/1/21/3392e861-03c7-47f8-8c8b-910914276e2d.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/1/21/18e74515-432f-4af6-92f6-f7881f4ff90f.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "case-hp-iphone-13-13-pro-13-pro-max-softcase-magsafe-premium-black-iphone-13pro-5b535",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 3",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 2 Nov",
    "toko": {
      "username": "delvinstore",
      "toko_name": "Delvin@store",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2024/1/21/c5e67e1d-e977-43d7-910a-da747b26edff.jpg"
    }
  },
  {
    "product_name": "Casing hp Samsung J2 Prime J7 J7Prime J72017 J7Pro Maple Lea Softcase - Hijau terang, J2 Prime",
    "price": "Rp25.000",
    "stock": "993",
    "description": "Selamat datang di toko kami ❤❤❤<br>Produk yang diperintahkan akan pabrik Indonesia dalam satu hari kerja, dan kami bergegas ke arahmu!<br><br>Kenapa kau memilih kami?<br>🤙KUALITAS TINGGI<br>Casing ponsel ini terbuat dari bahan softcase berkualitas tinggi serta, mudah dipasang &amp; dilepas, memberikan perlindungan terbaik untuk HP anda<br><br>🤙PENGIRIMAN LOKAL<br>Pengiriman langsung dari pabrik Indonesia untuk memastikan pengiriman barang yang cepat dan aman ke tangan Anda.<br><br>🤙LAYANAN YANG BAIK<br>Jika Anda memiliki pertanyaan tentang produk, layanan pelanggan akan meminta Anda online untuk layanan<br><br>Deskripsi Produk:<br>🤙💯Baru dan Dapat Dihapus.<br>🤙Desain pola kotak mode membuat ponsel Anda tetap pribadi dan terasa nyaman.<br>🤙Penyerap goncangan, bahan anti pecah.<br>🤙Lanyard menambahkan perlindungan terhadap jatuh yang tidak disengaja.<br>🤙Semua fungsi berfungsi dengan baik bahkan ponsel Anda dalam kasus ini.<br><br>Fitur:<br>Bahan: Terbuat dari Silikon Berkualitas Tinggi + TPU<br>Isi paket:<br>1 x Kasus Pelapisan + 1 x Lanyard",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/5/17/0de43673-dd0f-44ba-ad82-b5be3e5a71c2.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/5/17/fae74b8d-155d-4468-b792-2317bd6ffec2.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/5/17/39cf8a8d-0fdc-4fc9-9904-8027854b3426.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/5/17/b6d23f0b-81d9-4118-9f2e-aadee36a2cd5.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/5/17/fd7e6745-8998-464f-b988-22643342a841.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-samsung-j2-prime-j7-j7prime-j72017-j7pro-maple-lea-softcase-hijau-terang-j2-prime-31ba",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.7",
    "sold": "Terjual 100+",
    "country": "Kab. Tangerang",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "shopeper",
      "toko_name": "Trendi Case",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/4/29/1df532de-5202-4b3e-8db7-a42d37a779de.jpg"
    }
  },
  {
    "product_name": "Casing Hp Samsung A33 A53 5G Soft Case Plating Standing Astronot - Blue, Samsung A53 5G",
    "price": "Rp35.000",
    "stock": "984",
    "description": "Soft Case Import Casing Plating Dan Dudukan Handphone Astronot Premium<br><br>Ready Warna<br><br>- Black / Hitam<br>- Blue / Biru<br>- Purple / Ungu<br><br>Packing Menggunakan Kardus &amp; Bubble Warp<br><br>HIGH QUALITY AND PREMIUM DESIGN<br>- Case tidak hanya didesain dengan sangat elegant dan terlihat sangat mewah saat digunakan, namun case ini juga melindungi hp dengan baik dari benturan ataupun goresan<br>- Stand function (bisa diberdiriin)<br>- Menutupi sempurna bagian body samping dan belakang HP<br>- Mudah untuk dipasang dan dilepaskan<br>- Slim &amp; Fit<br>- Memberikan kesan mewah<br>- Bahan Berkualitas dan nyaman<br><br>Terima kasih..",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/6/25/dbf1b19b-407e-4559-9d5d-c247dc500d8e.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/6/25/2c856a95-dd01-4148-9dba-317681726147.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/6/25/2999e1ee-cf95-4b28-a439-7a63fc17236a.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/6/25/4ac4c7e6-0d0a-41a1-847d-e6a89fb42ce0.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-samsung-a33-a53-5g-soft-case-plating-standing-astronot-blue-samsung-a53-5g-22061",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.9",
    "sold": "Terjual 100+",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "fahryonemore",
      "toko_name": "FahryOneMore",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/9/7/3bb0f3ba-2d30-424a-885a-b4599e243647.jpg"
    }
  },
  {
    "product_name": "Case Samsung A51 A71 A7 2018 A750 casing hp cover tpu slim CALFSKIN - Black, A71",
    "price": "Rp19.000",
    "stock": "17",
    "description": "READY UNTUK :<br><br>- Samsung A12/M12<br>- Samsung A13 4G/5G/A04s<br>- Samsung A20/A30<br>- Samsung A22 4G/M22/M32<br>- Samsung A23 4G/5G<br>- Samsung A30s/A50/A50s<br>- Samsung A32 4G<br>- Samsung A33 5G<br>- Samsung A51<br>- Samsung A52/A52s 5G<br>- Samsung A53 5G<br>- Samsung A7 2018/A750<br>- Samsung A71<br>- Samsung A72 4G/5G<br>- Samsung A73 5G<br>- Samsung S7 Edge<br>- Samsung S7 Flat<br>- Samsung S8<br>- Samsung S8 Plus<br>- Samsung S10 Plus<br>- Samsung S20 Ultra<br>- Samsung S21 Ultra<br>- Samsung S22 Plus<br>- Samsung S22 Ultra<br>- Samsung Note 10<br>- Samsung Note 10 Plus<br>- Samsung Note 20<br>- Samsung Note 20 Ultra<br>- Samsung Note 8<br>- Samsung Note 9<br><br>Warna : Black, Brown, Green, Grey<br><br><br>INFO PENTING WAJIB !!!<br><br>1. Setiap pembukaan packing WAJIB MELAKUKAN VIDEO UNBOXING PAKET, jika produk rusak atau jumlahnya kurang akan divalidasi melalui video, jika tidak ada bukti video maka dianggap TIDAK SAH untuk komplain<br><br>2. PERHATIAN : KAMI TIDAK MENERIMA KOMPLAIN untuk produk rusak / pecah pada saat diterima ataupun rusak pada saat pengiriman. Jadi mohon kerjasamanya untuk membeli KARDUS PACKING PENGAMAN TAMBAHAN untuk meminimalisir terjadinya kerusakan pada pengiriman.<br><br>HARAP BACA SEBELUM BELI, PENTING !!!<br>- Semua Pesanan akan dikirim SESUAI VARIASI yang dipilih pembeli<br>- Tidak menerima request variasi via chat, catatan, ataupun hal lainnya<br>- MEMBELI = SETUJU<br><br>KAPAN PESANAN DIKIRIM ?<br>1. SENIN - JUMAT Order yang masuk sebelum Jam 15.00 akan dikirim pada hari itu juga, KECUALI Sameday &amp; Instant maks. jam 14.00<br>2. SABTU maks Jam 11.00 akan dikirim pada hari itu juga<br>3. MINGGU dan TANGGAL MERAH LIBUR !!!<br><br>DESKRIPSI PRODUK :<br>- Bahan Vintage Leather lembut dan nyaman + TPU Silikon<br>- Full Lens Protection, melindungi bagian camera dari goresan<br>- Anti-Drop Protection, melindungi dari benturan pada saat terjatuh<br>- Ada lubang tali untuk gantungan<br>- Mudah dipasang dan dilepaskan<br>- Lentur dan Flexible tidak merusak body HP",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2022/12/29/e9b69a7a-67b9-42c4-9eb0-b5772b618f9e.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/6/30/4c1db265-3403-405d-9452-3a31e2afe8cb.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/6/30/c6c50cc6-bd14-4b84-a769-5d53fa381ad9.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/6/30/3c2e27ab-4432-4fbc-8aa3-8a5266f69d17.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "case-samsung-a51-a71-a7-2018-a750-casing-hp-cover-tpu-slim-calfskin-black-a71",
    "discount": "53%",
    "discount_price": "Rp40.000",
    "rating": "4.9",
    "sold": "Terjual 100+",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "casepedias",
      "toko_name": "Case Pedia",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2017/1/5/1042576/1042576_4d8ee5d1-ccf0-4adc-aac9-5fbf9eeadc99.jpg"
    }
  },
  {
    "product_name": "Soft Case Samsung A34 A54 5G Clear Case TPU Silicon Casing Handphone - Samsung A34",
    "price": "Rp17.000",
    "stock": "871",
    "description": "NEW ARRIVAL..<br><br>READY STOCK WARNA CLEAR<br>====================================================<br><br>FEATURES<br>~Case Berbahan Thermoplastic Euthrane (TPU) Berkualitas.<br>~Ketinggian Case 0.5mm Pada Bagian Depan Untuk Mencegah Layar Utama Bersentuhan Langsung Pada Medan Datar.<br>~Pada Bagian Dalam Case Memiliki Texture Matte Agar Smartphone Anda Tidak Terlepas Ketika Terjadi Benturan.<br>~Pada Bagian Kamera, Case Memiliki Ketinggian Lebih Tinggi Agar Lensa Kamera Tidak Lecet.<br>~Tombol Yang Mudah Di Kontrol Karena Berbahan Soft.<br>~Bahan Tidak Gampang Menguning<br><br>Terima Kasih",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/3/22/38201301-b231-493a-813a-cf1bc603c834.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/3/22/45726b44-abea-4a28-b9d9-16d22d34dbde.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/3/22/4e1ffdf8-c7db-4b0a-a87d-d08dcc32031e.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/3/22/00cf8a69-7963-4db4-9f5d-06301405cfb9.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/3/22/8b760f3b-860c-4ea6-8369-a8226b8afa6e.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "soft-case-samsung-a34-a54-5g-clear-case-tpu-silicon-casing-handphone-samsung-a34",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 250+",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "fahryonemore",
      "toko_name": "FahryOneMore",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/9/7/3bb0f3ba-2d30-424a-885a-b4599e243647.jpg"
    }
  },
  {
    "product_name": "Casing Hp Oppo Reno 6 4G Soft Case Plating Standing Astronot - Blue",
    "price": "Rp35.000",
    "stock": "490",
    "description": "Soft Case Import Casing Plating Dan Dudukan Handphone Astronot Premium<br><br>Ready Warna<br><br>- Black / Hitam<br>- Blue / Biru<br>- Purple / Ungu<br><br>Packing Menggunakan Kardus &amp; Bubble Warp<br><br>HIGH QUALITY AND PREMIUM DESIGN<br>- Case tidak hanya didesain dengan sangat elegant dan terlihat sangat mewah saat digunakan, namun case ini juga melindungi hp dengan baik dari benturan ataupun goresan<br>- Stand function (bisa diberdiriin)<br>- Menutupi sempurna bagian body samping dan belakang HP<br>- Mudah untuk dipasang dan dilepaskan<br>- Slim &amp; Fit<br>- Memberikan kesan mewah<br>- Bahan Berkualitas dan nyaman<br><br>Terima kasih..",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/8/c9e33d21-018b-4008-ad0b-47872ddc66ce.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/8/5fa17f6a-f7fe-4aec-ae2d-08c433a837af.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/8/13721109-3340-4a58-953c-b17826b4d697.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/8/063dbe14-8c28-4c75-b09e-b08ffb2cb1bd.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-oppo-reno-6-4g-soft-case-plating-standing-astronot-blue-b63b8",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.9",
    "sold": "Terjual 60+",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "fahryonemore",
      "toko_name": "FahryOneMore",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/9/7/3bb0f3ba-2d30-424a-885a-b4599e243647.jpg"
    }
  },
  {
    "product_name": "al-Masjid an-Nabawi (Dome) - Case Hp Kaligrafi Nama (Tersedia Semua Tipe HP) - Ultimatte Black",
    "price": "Rp109.000",
    "stock": "99",
    "description": "AhviSunday&nbsp;adalah&nbsp;brand&nbsp;Custom&nbsp;Kaligrafi&nbsp;terlaris&nbsp;Di&nbsp;Indonesia.&nbsp;Ribuan&nbsp;pesanan&nbsp;untuk&nbsp;case&nbsp;HP&nbsp;sudah&nbsp;dipercayakan&nbsp;ke&nbsp;AhviSunday.<br><br>🔥&nbsp;Tersedia&nbsp;semua&nbsp;tipe&nbsp;HP&nbsp;:&nbsp;Vivo,&nbsp;Oppo,&nbsp;Xiaomi,&nbsp;Lenovo,&nbsp;Iphone,&nbsp;Samsung&nbsp;dll<br>🔥&nbsp;Kualitas&nbsp;Mesin&nbsp;UV&nbsp;asli&nbsp;Jepang<br>🔥&nbsp;Hasil&nbsp;cetak&nbsp;jelas&nbsp;dan&nbsp;tidak&nbsp;luntur<br>🔥&nbsp;Bahan&nbsp;ringan&nbsp;dalam&nbsp;genggaman<br>🔥&nbsp;Finishing&nbsp;Doff&nbsp;/&nbsp;Glossy<br><br>==================================================<br><br>Informasi&nbsp;Bahan&nbsp;:<br><br>Ultimatte&nbsp;Black<br>✔️&nbsp;Bahan&nbsp;dasar&nbsp;silikon<br>✔️&nbsp;Cetak&nbsp;bagian&nbsp;permukaan&nbsp;belakang<br>✔️&nbsp;Lentur&nbsp;dan&nbsp;fit&nbsp;dengan&nbsp;smartphone<br>✔️&nbsp;Dasar&nbsp;Case&nbsp;berwarna&nbsp;hitam<br><br>Clear&nbsp;Case<br>✔️&nbsp;Bahan&nbsp;dasar&nbsp;silikon&nbsp;bening<br>✔️&nbsp;Cetak&nbsp;bagian&nbsp;permukaan&nbsp;belakang<br>✔️&nbsp;Lentur&nbsp;dan&nbsp;fit&nbsp;dengan&nbsp;smartphone<br>✔️&nbsp;Elastis&nbsp;bisa&nbsp;kembali&nbsp;ke&nbsp;bentuk&nbsp;semula<br>✔️&nbsp;Ada&nbsp;sisi&nbsp;bumper&nbsp;di&nbsp;sudut&nbsp;untuk&nbsp;meminimalisir&nbsp;retak&nbsp;apabila&nbsp;terjatuh<br><br>Slim&nbsp;Clear&nbsp;Case<br>✔️&nbsp;Bahan&nbsp;dasar&nbsp;silikon&nbsp;bening<br>✔️&nbsp;Cetak&nbsp;bagian&nbsp;permukaan&nbsp;belakang<br>✔️&nbsp;Lentur&nbsp;dan&nbsp;fit&nbsp;dengan&nbsp;smartphone<br>✔️&nbsp;Elastis&nbsp;bisa&nbsp;kembali&nbsp;ke&nbsp;bentuk&nbsp;semula<br><br>Exclusive&nbsp;Case<br>✔️&nbsp;Anti&nbsp;Gores,&nbsp;Anti&nbsp;Api&nbsp;dan&nbsp;Anti&nbsp;Noda<br>✔️&nbsp;Full&nbsp;Glass&nbsp;/&nbsp;plat&nbsp;besi&nbsp;dengan&nbsp;Coating&nbsp;TPU<br>✔️&nbsp;Print&nbsp;pada&nbsp;permukaan<br>✔️&nbsp;Menggunakan&nbsp;sistem&nbsp;cetak&nbsp;Transfer&nbsp;Printing&nbsp;Film<br>✔️&nbsp;Bevel&nbsp;di&nbsp;ujung&nbsp;case&nbsp;agar&nbsp;glass&nbsp;tidak&nbsp;tajam<br><br>Fullprint&nbsp;Hardcase<br>✔️&nbsp;Bahan&nbsp;Polycarbonate<br>✔️&nbsp;Full&nbsp;print&nbsp;/&nbsp;cetak&nbsp;hingga&nbsp;sisi&nbsp;samping&nbsp;case<br>✔️&nbsp;Hasil&nbsp;jelas&nbsp;dan&nbsp;tidak&nbsp;luntur<br>✔️&nbsp;Atas&nbsp;dan&nbsp;bawah&nbsp;case&nbsp;terbuka<br>✔️&nbsp;Bahan&nbsp;ringan&nbsp;di&nbsp;dalam&nbsp;genggaman&nbsp;walau&nbsp;bahan&nbsp;keras<br><br>==================================================<br><br>Cara&nbsp;Pemesanan&nbsp;:<br><br>⚠️&nbsp;PENTING!!!&nbsp;Tanyakan&nbsp;ketersediaan&nbsp;stok&nbsp;case&nbsp;terlebih&nbsp;dahulu<br>✔️&nbsp;Pilih&nbsp;jenis&nbsp;case<br>✔️&nbsp;Masukkan&nbsp;gambar&nbsp;desain&nbsp;pada&nbsp;keterangan&nbsp;/&nbsp;chat&nbsp;via&nbsp;tokopedia<br>✔️&nbsp;Mohon&nbsp;cantumkan&nbsp;nama&nbsp;kaligrafi&nbsp;beserta&nbsp;tipe&nbsp;hpnya&nbsp;dicatatan&nbsp;pembelian<br><br>FREE&nbsp;Garansi&nbsp;/&nbsp;Retur&nbsp;Jika&nbsp;diterima&nbsp;kondisi&nbsp;rusak",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/db79be60-bf19-43fc-bfe5-455f89f241aa.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/eaeb4e80-6bdd-4ecf-8b1f-6d9f6f95d9ea.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/fac92a4b-2288-41a7-b39b-6db402d829f0.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/db79be60-bf19-43fc-bfe5-455f89f241aa.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/20/50d6fde3-d9b4-4885-a409-5bfae129c215.png.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "al-masjid-an-nabawi-dome-case-hp-kaligrafi-nama-tersedia-semua-tipe-hp-ultimatte-black-e41dc",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 1",
    "country": "Kab. Cianjur",
    "estimation": "Estimasi tiba 14 - 17 Nov",
    "toko": {
      "username": "ahvisunday",
      "toko_name": "AhviSunday",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2022/6/7/a74cd2c3-c2fd-40fa-a651-72321b1b6bf4.jpg"
    }
  },
  {
    "product_name": "Softcase Anti Crack Elegant Casing Hp Redmi Note 10 Pro Free TG",
    "price": "Rp14.449",
    "stock": "1455",
    "description": "- Produk Tidak Semua Persis Seperti Gambar<br>- Barang yang anda bisa lihat ready stock dan bisa langsung dipesan.<br>- Jika barang yang dibeli ternyata habis, kami akan menggantikan dengan uang cash atau ditukar dengan barang lain yang harganya sama.<br>- Barang yang dijual lengkap sesuai dengan isi foto.<br>- Setiap orderan yang kami terima, kami usahakan proses paling lambat 2 X 24 jam.<br>- Tidak menerima Retur/Komplen barang yang rusak/cacat karena kami Sudah Mnegecek Brang sebelum Dikirim Dengan Kondisi Baik<br>- Untuk orderan yang masuk di hari Minggu akan kami proses di hari SENIN.<br>- CHECK dan PASTIKAN lebih teliti apakah alamat / nama / nomor handphone sudah sesuai, karena kami tidak menerima perubahan apapun jika sudah tercantum didalam invoice kami.<br>- Jika ada PERTANYAAN bisa ditanyakan di kolom diskusi atau pesan.<br>- Nomor Resi akan di input secepatnya.<br>Terima kasih telah mampir di toko kami, semoga bisa bekerja sama, salam sukses dan selamat berbelanja...",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/1/2/f43be0fb-6877-438a-b010-d7403d1246f5.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/1/2/e5c9da9d-5bb1-4fdb-831a-0888dd306462.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/1/2/1c9c138c-a0ea-47e2-a4ee-653c21566cba.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "softcase-anti-crack-elegant-casing-hp-redmi-note-10-pro-free-tg",
    "discount": "4%",
    "discount_price": "Rp15.000",
    "rating": "5",
    "sold": "Terjual 40+",
    "country": "Kota Bekasi",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "asmanofficial",
      "toko_name": "Asman official",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2019/9/29/6195551/6195551_9dc2fb26-14b4-4f11-905a-ffe8e262e71a.jpg"
    }
  },
  {
    "product_name": "VIVO V20 V23 V23E V25 V25E V27 V27E V29 V29E V30 V30 PRO HOLOGRAM CASING VIRAL KESING HP CESING HANDPHONE SOFT HARD CASE - Hitam, VIVO V20",
    "price": "Rp15.000",
    "stock": "1.000",
    "description": "KELEBIHAN PRODUK :<br><br>Bahan material : Hybrid campuran softcase dan hardcase<br>- Bahan TPU Premium Bertekstur Lembut<br>- Nyaman di Genggam Tidak Licin<br>- Melindungi Dari Benturan Dengan Sempurna<br>- Anti Scratch (Tidak Mudah Tergores)<br>- Tidak Menganggu Sinyal Ponsel<br>- Tidak menimbulkan baret pada bodi ponsel<br>- Anti Fingerprint apabila jari kita berminyak atau kotor walaupun terkena Case akan sangat mudah untuk dibersihkan<br><br>MEMBELI = SETUJU<br>tidak menerima returan jika tidak ada video unboxing<br>#Liang's Store",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/5d667e4d-2f6a-4e30-a6e9-4d14b6c22b6a.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/495837b9-fcc0-4a0a-b6d5-9222ae7df253.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/17a66ece-fe40-4bea-adb1-247bbe128618.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/0c4c64db-9da3-4a94-a622-572e4c4f18b9.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/5d667e4d-2f6a-4e30-a6e9-4d14b6c22b6a.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/9a008fe9-6eea-4802-a24b-f65a96e3e601.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/495837b9-fcc0-4a0a-b6d5-9222ae7df253.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/17a66ece-fe40-4bea-adb1-247bbe128618.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/0c4c64db-9da3-4a94-a622-572e4c4f18b9.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "vivo-v20-v23-v23e-v25-v25e-v27-v27e-v29-v29e-v30-v30-pro-hologram-casing-viral-kesing-hp-cesing-handphone-soft-hard-case-hitam-vivo-v20-e5172",
    "discount": "null",
    "discount_price": "null",
    "rating": "null",
    "sold": "null",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "liangsstore",
      "toko_name": "Liang's Store",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2022/10/11/39000dc2-26cd-4750-988d-968d4c5b8bed.jpg"
    }
  },
  {
    "product_name": "Case hp iphone 15 14 plus pro max premium casing keren - Hitam, Ip 14",
    "price": "Rp75.000",
    "stock": "100",
    "description": "wellcome sanelvisacc<br><br>casing magsafe pc<br><br>tersedia warna<br>~ Black<br>~ Gold<br>~ Rose Gold<br>~ Deep purple<br>~ Biru<br><br>happy shopping🙂",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/28/c0a3d1f5-b9dc-4d5b-b873-a48b96a4871c.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/28/6a28a15e-9d26-4592-8ad5-6354e6bffd9e.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/28/56992d28-6852-44f6-8c41-47747f33ad33.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/28/c750d04a-a8db-4299-9b8b-5a0b33d04fae.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/8/28/e7358f34-a1cf-4c06-b92e-e91b1585dfb3.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "case-hp-iphone-15-14-plus-pro-max-premium-casing-keren-hitam-ip-14-743e4",
    "discount": "null",
    "discount_price": "null",
    "rating": "null",
    "sold": "null",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "sanelvisacc",
      "toko_name": "Sanelvisacc",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2022/6/23/8422baec-4e21-4bd7-8421-ba75ade44750.jpg"
    }
  },
  {
    "product_name": "Colta Aluminium Magsafe Phone Case Casing HP iPhone Wireless Charger Powerbank - Brown, iPhone 14PRO",
    "price": "Rp259.000",
    "stock": "13",
    "description": "𝐀𝐥𝐮𝐦𝐢𝐧𝐢𝐮𝐦 𝐌𝐚𝐠𝐬𝐚𝐟𝐞 𝐂𝐚𝐬𝐞<br><br>𝐒𝐩𝐞𝐜𝐢𝐟𝐢𝐜𝐚𝐭𝐢𝐨𝐧 :<br>- Model : CASE-ALU<br>- Warna : Leather Black/Leather Brown<br>- Screen Lips : 1.5MM<br>- Camera Lips : 1.55MM<br>- Back Lips : 0.52MM<br><br>𝐁𝐚𝐡𝐚𝐧 : 𝐀𝐥𝐮𝐦𝐢𝐧𝐢𝐮𝐦 &amp; 𝐋𝐞𝐚𝐭𝐡𝐞𝐫 (𝐇𝐲𝐛𝐫𝐢𝐝 𝐂𝐚𝐬𝐞)<br>- Hard Aluminium Bumper<br>- Hard Aluminium + Leather Back<br>- Neodymium Permanent Magnet<br><br>- Berat : 184 gram<br>- Support Wireless Charging<br>- Magnetic Technology<br><br>⚠️⚠️⚠️<br>𝐖𝐚𝐫𝐫𝐚𝐧𝐭𝐲 𝐏𝐨𝐥𝐢𝐜𝐲:<br>1. Unboxing Video required.<br>2. If damage occured on our products caused by courier, within 3 days after receiving the package please do inform us the evidence in video and photos. We will assign you to get the new product.<br>(⚠️Attention: More than 3 days, no warranty available)<br><br>𝐏𝐚𝐜𝐤𝐚𝐠𝐞 𝐈𝐧𝐜𝐥𝐮𝐝𝐞𝐬:<br>1x Aluminium Magsafe Case<br><br><br>𝐍𝐨𝐭𝐞𝐬 :<br>1. Harap pastikan pesanan sudah sesuai sebelum check out.<br>2. Produk tidak dapat ditukar size / warna / model<br>3. Untuk klaim produk tidak lengkap harap sertakan video unboxing paket.<br>4. Setelah barang sampai mohon dicek kelengkapan produk dan pastikan barang diterima dengan kondisi baik sebelum klik \"Pesanan Diterima\".<br><br>𝐏𝐞𝐫𝐢𝐧𝐠𝐚𝐭𝐚𝐧 : Sebelum memberi penilaian diharapkan untuk menghubungi kami terlebih dahulu, kami akan memberi layanan terbaik kepada anda, penilaian bintang anda sangat berarti untuk kami, Terima kasih~",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/4/23/16383f11-5927-4cd9-b00c-33a89e908bbc.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/4/20/ef681cae-383b-4312-8573-140dc1718f48.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/4/20/29d930db-34ca-48e2-98c8-e60a1470311f.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/4/20/eef45685-687d-4288-b105-005ba247494f.png.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "colta-aluminium-magsafe-phone-case-casing-hp-iphone-wireless-charger-powerbank-brown-iphone-14pro-e3a52",
    "discount": "35%",
    "discount_price": "Rp399.000",
    "rating": "5",
    "sold": "Terjual 22",
    "country": "Kota Surabaya",
    "estimation": "Estimasi tiba 31 Okt - 4 Nov",
    "toko": {
      "username": "colta",
      "toko_name": "Colta Official Store",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2022/3/29/8ee9ba14-9e8e-47ed-b008-bb97bd2994b5.png"
    }
  },
  {
    "product_name": "Soft Case OPPO A71 A83 Kesing Casing Silicon Silikon Kondom Sopkes HP - Rose, A71",
    "price": "Rp7.900",
    "stock": "4",
    "description": "100% Gambar Asli / Real Pictures<br>100% Gambar Asli / Real Pictures<br><br>- Terbuat dari bahan Silicon Premium<br>- Terdapat pelindung dibagian kamera<br>- Tidak mudah kotor dan berubah warna<br>- Permukaan lembut dan tidak licin<br>- Tahan benturan, Kuat dan Awet<br>- Elastis dan lentur tapi tidak mudah melar!!<br><br> &gt;&gt;Tersedia 7 warna :<br>- Hitam<br>- Biru<br>- Coklat<br>- Hijau<br>- Rose (Pink Kecoklatan)<br>- Maroon<br>- Broken White<br><br>**sertakan pilihan warna cadangan pada catatan<br>**jika warna yg dipilih kosong akan kami konfirmasi",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/10/22/7ed72b5d-44f1-4c5f-a9c6-a08c68ca34b8.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/10/22/862eaa7c-bc76-447a-a71e-e3b6766749a0.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/10/22/953f9c4f-8055-4b37-a217-d611e79a2201.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/10/22/3cc80f3d-49f3-45e5-b677-7497b505eea2.png.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "soft-case-oppo-a71-a83-kesing-casing-silicon-silikon-kondom-sopkes-hp-rose-a71",
    "discount": "68%",
    "discount_price": "Rp24.900",
    "rating": "4.9",
    "sold": "Terjual 250+",
    "country": "Kota Bekasi",
    "estimation": "Estimasi tiba besok - 2 Nov",
    "toko": {
      "username": "rajaapasaja",
      "toko_name": "Raja Apa Saja",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/4/29/41f8fb71-cdd2-475f-9cfd-beea726032dc.png"
    }
  },
  {
    "product_name": "Casing Hp Xiaomi POCO M3 Case Handphone Softcase Hardcase - 114 - 1, POCO M3",
    "price": "Rp19.500",
    "stock": "233",
    "description": "CASING CASE PREMIUM TERSEDIA UNTUK INFINIX # OPPO # REALME # SAMSUNG # VIVO # IPHONE # ADVAN #XIAOMI #ASUS JIKA INGIN ORDER TIPE LAIN TULIS DI CATATAN YA KAK JADI KAKAK TINGGAL PILIH GAMBAR MANA YANG KAKAK SUKA . LALU UTNTUK TIPE HP NYA DI CANTUMIN KE CATATAN ATAU BISA LANGSUNG CANTUMIN TIPE NYA DI CHAT YA.. CONTOH ORDER (NO 1 UNTUK OPPO Y12S)",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/2/24/112800df-dc58-4e52-a4dd-626f589e9659.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/2/24/a6777ce9-4f51-4c1f-b68a-c12f93c430b6.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/2/24/c80fa362-84cc-41ce-a6d1-45f612742400.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/2/24/ffce94bd-b4a5-4ed9-8379-a1c1c64e94e0.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/2/24/112800df-dc58-4e52-a4dd-626f589e9659.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-xiaomi-poco-m3-case-handphone-softcase-hardcase-114-1-poco-m3-854b6",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.8",
    "sold": "Terjual 100+",
    "country": "Kota Bogor",
    "estimation": "Estimasi tiba besok - 2 Nov",
    "toko": {
      "username": "intristore",
      "toko_name": "intristore",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2021/9/8/3677650a-38dd-4c90-a675-9e651b3cd681.jpg"
    }
  },
  {
    "product_name": "Casing hp Samsung A32 A33 5G A52 A52S A53 A72 A73 5G Case Bunga Seni - Gaya 1, Samsung A32 5G",
    "price": "Rp15.000",
    "stock": "995",
    "description": "Selamat datang di toko kami ❤❤❤<br>Produk yang diperintahkan akan pabrik Indonesia dalam satu hari kerja, dan kami bergegas ke arahmu!<br><br>Kenapa kau memilih kami?<br>🤙KUALITAS TINGGI<br>Casing ponsel ini terbuat dari bahan softcase berkualitas tinggi serta, mudah dipasang &amp; dilepas, memberikan perlindungan terbaik untuk HP anda<br><br>🤙PENGIRIMAN LOKAL<br>Pengiriman langsung dari pabrik Indonesia untuk memastikan pengiriman barang yang cepat dan aman ke tangan Anda.<br><br>🤙LAYANAN YANG BAIK<br>Jika Anda memiliki pertanyaan tentang produk, layanan pelanggan akan meminta Anda online untuk layanan<br><br>Deskripsi Produk:<br>🤙💯Brand baru dan dapat dihapus.<br>🤙Desain wavy mode menjaga ponselmu pribadi dan merasa baik.<br>🤙Membebaskan pola yang sama hadiah pemilik kembali.<br>🤙Pemilik belakang mendukung lokasi telepon 360 derajat ketika digunakan.<br>🤙Penyerap kejutan, materi silikon lembut yang tak bisa hancur.<br>🤙Desain lensa unik untuk menghindari goresan kamera.<br>🤙Lubang kamera yang tebal menambah perlindungan terhadap titik kecelakaan.<br>🤙Bahkan jika ponselmu berada di kondisi ini, semua fungsi akan bekerja dengan baik.<br><br>Fitur:<br>Bahan: Dibuat dari TPU Kualitas Tinggi<br>Isi paket:<br>1 x Kasing",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/5/19/47590a76-8a4f-4c77-87a4-2f7f0d89f975.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/5/19/b2a89418-02b4-40fb-b215-127bdb342610.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/5/19/9a9b60e7-2c2e-4b9a-b267-286ca9e3c0cc.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/5/19/6ab0131f-2830-4069-8501-44ccef8f8cb8.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/5/19/eb303635-bc4d-4577-a6bb-c97d97610955.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-samsung-a32-a33-5g-a52-a52s-a53-a72-a73-5g-case-bunga-seni-gaya-1-samsung-a32-5g",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.9",
    "sold": "Terjual 40+",
    "country": "Kab. Tangerang",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "shopeper",
      "toko_name": "Trendi Case",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/4/29/1df532de-5202-4b3e-8db7-a42d37a779de.jpg"
    }
  },
  {
    "product_name": "Xiaomi Redmi 13C - Poco C65 Case Softcase WINDOW STAND MACARON Case Casing Hp Redmi 13C - C65 - Hijau, Redmi 13C",
    "price": "Rp22.500",
    "stock": "996",
    "description": "NEW DESIGN - NEW MODEL<br>WINDOW STAND MACARON<br><br>- PREMIUM QUALITY SILICONE<br>- LEMBUT - TIDAK LICIN - KESET <br>- FRAME KAMERA BISA DI JADIKAN STAND / BRACKET <br>- PRESISI - SAFETY - PROTECT YOUR PHONE<br><br>WAJIB MENCANTUMKAN WARNA CADANGAN KETIKA PEMBELIAN (TULIS DI CATATAN PEMBELIAN / CHAT KE TOKO)<br>TIDAK MENCANTUMKAN WARNA CADANGAN AKAN KAMI RANDOM BILAMANA WARNA YANG DIPILIH KOSONG<br>",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/27/41d38b1d-748d-46d9-a5fa-fe8476324a75.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/27/1cf9eda9-f225-4c94-b5ff-f113a4a01276.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/27/d44e6953-58bc-40ef-995b-51bd9ba5d1f5.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/27/7f88c093-7c44-4d73-9e2e-17cb019e30fa.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/27/7973d950-5061-465c-aa99-ffe5aed9d16d.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "xiaomi-redmi-13c-poco-c65-case-softcase-window-stand-macaron-case-casing-hp-redmi-13c-c65-hijau-redmi-13c-ec363",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.9",
    "sold": "Terjual 60+",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 2 Nov",
    "toko": {
      "username": "drawacc",
      "toko_name": "Drawacc",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2024/7/2/1a9dd5a8-be87-4e4d-a19c-e3dc1228efab.jpg"
    }
  },
  {
    "product_name": "Softcase Xiaomi Redmi A2 A1 10 10A 10C 10 5G Seru kupu -kupu Casing hp - Hijau terang, Redmi 10",
    "price": "Rp25.000",
    "stock": "995",
    "description": "Selamat datang di toko kami ❤❤❤<br>Produk yang diperintahkan akan pabrik Indonesia dalam satu hari kerja, dan kami bergegas ke arahmu!<br><br>Kenapa kau memilih kami?<br>🤙KUALITAS TINGGI<br>Casing ponsel ini terbuat dari bahan Casing hp berkualitas tinggi serta, mudah dipasang &amp; dilepas, memberikan perlindungan terbaik untuk HP anda<br><br>🤙PENGIRIMAN LOKAL<br>Pengiriman langsung dari pabrik Indonesia untuk memastikan pengiriman barang yang cepat dan aman ke tangan Anda.<br><br>🤙LAYANAN YANG BAIK<br>Jika Anda memiliki pertanyaan tentang produk, layanan pelanggan akan meminta Anda online untuk layanan<br><br>Deskripsi Produk:<br>🤙💯Baru dan Dapat Dihapus.<br>🤙Desain pola kotak mode membuat ponsel Anda tetap pribadi dan terasa nyaman.<br>🤙Penyerap goncangan, bahan anti pecah.<br>🤙Lanyard menambahkan perlindungan terhadap jatuh yang tidak disengaja.<br>🤙Semua fungsi berfungsi dengan baik bahkan ponsel Anda dalam kasus ini.<br><br>Fitur:<br>Bahan: Terbuat dari Silikon Berkualitas Tinggi + TPU<br>Isi paket:<br>1 x Softcase elektroplating + 1 x Lanyard",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/6/8/a4ec593c-9669-484e-a563-88578218a218.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/6/8/7b2d22db-48d4-487a-892a-59a289f4199e.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/6/8/2a052a95-e0c0-477c-90f3-23581deeed2d.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/6/8/47ef3a1d-9718-4367-b4e9-6b5a06e8e4ee.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/6/8/b0d32745-ce46-4b2c-887f-11a4552aa942.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "softcase-xiaomi-redmi-a2-a1-10-10a-10c-10-5g-seru-kupu-kupu-casing-hp-hijau-terang-redmi-10",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 24",
    "country": "Kab. Tangerang",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "shopeper",
      "toko_name": "Trendi Case",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/4/29/1df532de-5202-4b3e-8db7-a42d37a779de.jpg"
    }
  },
  {
    "product_name": "Soft Case HP Samsung A21S A21 S Silicon Kesing Casing Cover Kondom - Coklat",
    "price": "Rp7.900",
    "stock": "14",
    "description": "100% Gambar Asli / Real Pictures<br>100% Gambar Asli / Real Pictures<br><br>- Terbuat dari bahan Silicon Premium<br>- Terdapat pelindung dibagian kamera<br>- Tidak mudah kotor dan berubah warna<br>- Permukaan lembut dan tidak licin<br>- Tahan benturan, Kuat dan Awet<br>- Elastis dan lentur tapi tidak mudah melar!!<br><br>&gt;&gt;Tersedia 7 warna :<br>- Hitam<br>- Biru<br>- Coklat<br>- Hijau<br>- Rose<br>- Maroon<br>- Broken White<br><br>**sertakan pilihan warna cadangan pada catatan<br>**jika warna yg dipilih kosong akan kami konfirmasi<br>",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/9/20/7456a53f-8824-4ec2-9faf-46b458367a6e.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/9/20/bdecfe97-9f55-4088-b3fb-49fd3ef8a3ee.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/9/20/1b3b2347-f2e0-47d5-9393-0b8d1b899136.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/9/20/b46c27fd-7216-4726-9a1e-c3f6b5c99769.png.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "soft-case-hp-samsung-a21s-a21-s-silicon-kesing-casing-cover-kondom-coklat",
    "discount": "68%",
    "discount_price": "Rp24.900",
    "rating": "4.9",
    "sold": "Terjual 250+",
    "country": "Kota Bekasi",
    "estimation": "Estimasi tiba besok - 2 Nov",
    "toko": {
      "username": "rajaapasaja",
      "toko_name": "Raja Apa Saja",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/4/29/41f8fb71-cdd2-475f-9cfd-beea726032dc.png"
    }
  },
  {
    "product_name": "Casing HP HONKAI STAR RAIL OPPO XIAOMI IPHONE SAMSUNG CUSTOM CASE - Softcase TPU",
    "price": "Rp95.000",
    "stock": "934",
    "description": "HONKAI STAR RAIL COLLECTION<br><br>🔥 Tersedia semua tipe HP : Vivo, Oppo, Xiaomi, Lenovo, Iphone, Samsung dll<br>🔥 Kualitas Mesin UV asli Jepang<br>🔥 Hasil cetak jelas dan tidak luntur<br>🔥 Bahan ringan dalam genggaman<br>🔥 Finishing Doff / Glossy<br><br>==================================================<br><br>Informasi Bahan :<br><br>Softcase TPU<br>✔️ Bahan dasar silikon<br>✔️ Cetak bagian permukaan belakang<br>✔️ Lentur dan fit dengan smartphone<br>✔️ Bahan dasar transparan<br>✔️ Bahan lebih tebal<br><br>Softcase Black Matte HG<br>✔️ Bahan dasar silikon<br>✔️ Cetak bagian permukaan belakang<br>✔️ Lentur dan fit dengan smartphone<br>✔️ Bahan dasar berwarna hitam<br>✔️ Bahan lebih tebal<br><br>Softcase Anticrack HG<br>✔️ Bahan dasar silikon<br>✔️ Cetak bagian permukaan belakang<br>✔️ Lentur dan fit dengan smartphone<br>✔️ Elastis bisa kembali ke bentuk semula<br>✔️ Untuk anticrack ada sisi bumper di sudut untuk meminimalisir retak apabila terjatuh<br>✔️ Bahan lebih tebal<br><br>Tempered Glass Case<br>✔️ Anti Gores, Anti Api dan Anti Noda<br>✔️ Full Glass dengan Coating TPU<br>✔️ Print pada permukaan Coating Glass<br>✔️ Menggunakan sistem cetak Transfer Printing Film<br>✔️ Terdapat Bevel di ujung case agar glass tidak tajam<br><br>Hardcase 3D<br>✔️ Bahan Polycarbonate<br>✔️ Full print / cetak hingga sisi samping case<br>✔️ Hasil jelas dan tidak luntur<br>✔️ Bahan ringan di dalam genggaman walau bahan keras<br>✔️ Sisi atas, bawah dan tombol terbuka agar mengikat ke smartphone<br><br>==================================================<br><br>Cara Pemesanan :<br><br>⚠️ PENTING!!! Tanyakan ketersediaan stok case terlebih dahulu<br>✔️ Pilih jenis case<br>✔️ Masukkan gambar desain pada keterangan / chat via tokopedia<br><br>Happy Shopping 🔥",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/9/4/8c38c333-928f-4def-871b-a4d6495a5f14.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/5/23/a318456a-9cef-41c8-938f-069a04cfbb75.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/9/4/b0eb91df-81cf-458e-8095-5a8af5cba9de.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/4/23/609d10db-7321-4cbf-8a48-6a5ee75be35e.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/1/26/e034951e-1f7a-423f-b8cc-7817ba675cad.png.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-honkai-star-rail-oppo-xiaomi-iphone-samsung-custom-case-softcase-tpu-d6110",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.9",
    "sold": "Terjual 500+",
    "country": "Kab. Cianjur",
    "estimation": "Estimasi tiba 15 - 19 Nov",
    "toko": {
      "username": "caseonstore",
      "toko_name": "Caseon Store",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2020/8/8/8034535/8034535_eb4f1c05-ee15-4329-8c62-ed40fa6dc586.png"
    }
  },
  {
    "product_name": "Paddy Premium Case Design Collaboration - Casing Custom Semua Tipe HP - V1V0",
    "price": "Rp39.000",
    "stock": "382",
    "description": "PENTING :<br><br>Pastikan kamu chat admin kami dulu untuk mengkonfirmasi ketersediaan tipe HP kamu. Jika sudah order dan case kosong, kita akan kirim case yang tersedia atau produk lain ya karena tidak ada pembatalan order di Paddy. Terimakasih<br><br>Kreasikan casing handphone kamu bersama Paddy Case! Ada berbagai jenis casing handphone dan design yang bisa kamu pilih sesuai keinginan kamu.<br><br>Kualitas printing Paddy menggunakan teknologi UV Print tercanggih dari Jepang sehingga hasil printnya tajam, super HD, anti gores, tidak luntur , bisa dicuci dan tahan lama.<br><br>Untuk cara order silahkan cek di gambar ya!<br><br>Jika ada pertanyaan silahkan hubungi admin kami via chat !<br><br>PASTIKAN ORDERAN KAMU SUDAH LENGKAP dan JELAS YA BESTIES!!<br>Jika kami menerima order yang tidak jelas (warna, design, Tipe HP dll), dan tidak ada balasan dari chat konfirmasi admin kami, maka kami akan mengirimkan design secara random ya.",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/6/11/d7ba1862-60a6-4049-85f7-b86473a099b9.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/5/6/3ddb5b7f-4281-45c5-ab42-e50763cd4923.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/1/31/8987344a-8dcb-45e1-b573-c9708bb863e8.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/1/31/7002a7d6-5f37-4bac-92d0-3c86098e54d9.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "paddy-premium-case-design-collaboration-casing-custom-semua-tipe-hp-v1v0",
    "discount": "61%",
    "discount_price": "Rp100.000",
    "rating": "4.9",
    "sold": "Terjual 750+",
    "country": "Kab. Bandung",
    "estimation": "Estimasi tiba 13 - 16 Nov",
    "toko": {
      "username": "paddyofficial",
      "toko_name": "Paddy Official",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2022/11/17/9bfe1115-2fea-4840-b89c-66d8c51e5b48.jpg"
    }
  },
  {
    "product_name": "TPU case Samsung A6 2018 - A6 Plus softcase casing hp cover ultra thin - A6",
    "price": "Rp12.000",
    "stock": "36",
    "description": "READY UNTUK :<br><br>- Samsung A6 2018<br>- Samsung A6 Plus<br><br>Warna : Black<br><br>HARAP BACA SEBELUM BELI, PENTING !!!<br>- Pesanan akan dikirim sesuai variasi yang dipilih<br>- Tidak menerima request variasi via chat<br>- Jika tidak ada variasi, harap tulis catatan pada pesanan<br>- Jika tidak mencantumkan catatan apapun, produk akan dikirim acak<br>- MEMBELI = SETUJU<br><br>KAPAN PESANAN DIKIRIM ?<br>1. SENIN - JUMAT Order yang masuk sebelum Jam 15.00 akan dikirim pada hari itu juga, KECUALI Sameday &amp; Instant maks. jam 14.00<br>2. SABTU maks Jam 11.00 akan dikirim pada hari itu juga<br>3. MINGGU dan TANGGAL MERAH LIBUR !!!<br><br>DESKRIPSI PRODUK :<br>- Bahan softcase TPU lentur<br>- Ketebalan 0.9mm sangat tipis<br>- Ultra Thin<br>- Tekstur matte / doff tidak licin<br>- Camera lens Protection (Proteksi Camera)<br>- Tidak membekas sidik jari (Anti Fingerprint)<br>- Tombol dan lubang-lubang presisi",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2020/8/6/07b44903-95de-409f-b68a-fe55f0e68c33.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2020/8/6/a879a0d7-cfc9-4b0d-8c2d-6853f447577b.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2020/8/6/75f3fdf7-39e6-4356-83a4-c7b1f02432f8.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2020/8/6/6010e614-cf62-4c74-ad40-2d09816ae3ac.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2020/8/6/18c8a353-e7bf-4249-b952-d5ff05448fa5.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "tpu-case-samsung-a6-2018-a6-plus-softcase-casing-hp-cover-ultra-thin-a6",
    "discount": "29%",
    "discount_price": "Rp17.000",
    "rating": "4.9",
    "sold": "Terjual 750+",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "casing-handphone",
      "toko_name": "Casing handphone murah",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2018/5/18/842194/842194_9c451296-efe8-4b91-b836-5525db13cf18.jpeg"
    }
  },
  {
    "product_name": "REALME NOTE 50 NARZO 20 30A 50A 50i PRIME HOLOGRAM CASING VIRAL KESING HP CESING HANDPHONE SOFT HARD CASE realmi - Hitam, REALME NOTE 50",
    "price": "Rp15.000",
    "stock": "1.000",
    "description": "KELEBIHAN PRODUK :<br><br>Bahan material : Hybrid campuran softcase dan hardcase<br>- Bahan TPU Premium Bertekstur Lembut<br>- Nyaman di Genggam Tidak Licin<br>- Melindungi Dari Benturan Dengan Sempurna<br>- Anti Scratch (Tidak Mudah Tergores)<br>- Tidak Menganggu Sinyal Ponsel<br>- Tidak menimbulkan baret pada bodi ponsel<br>- Anti Fingerprint apabila jari kita berminyak atau kotor walaupun terkena Case akan sangat mudah untuk dibersihkan<br><br>MEMBELI = SETUJU<br>tidak menerima returan jika tidak ada video unboxing<br>#Liang's Store",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/ec2662b1-b99a-4062-bf98-bada6be4b559.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/575db99f-74b7-433a-8eff-11548433145a.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/5da6a257-365d-465e-8b4a-fb0908db5d22.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/6bc256df-4934-4bab-b168-379c45769a4d.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "realme-note-50-narzo-20-30a-50a-50i-prime-hologram-casing-viral-kesing-hp-cesing-handphone-soft-hard-case-realmi-hitam-realme-note-50-561a7",
    "discount": "null",
    "discount_price": "null",
    "rating": "null",
    "sold": "null",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "liangsstore",
      "toko_name": "Liang's Store",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2022/10/11/39000dc2-26cd-4750-988d-968d4c5b8bed.jpg"
    }
  },
  {
    "product_name": "Baseus Cylinder Slip Cover Waterproof Bag Phone Case HP Anti Air Basah - Hitam",
    "price": "Rp79.000",
    "stock": "99,9 rb",
    "description": "Selamat datang ditoko kami, kami merupakan salah satu distributor resmi dari merk Baseus di Indonesia. Produk kami jamin bahwa 100% original.<br><br>THIS IS 100% ORIGINAL PRODUCT<br><br>1.5 lapisan segel tahan air<br><br>2.Peringkat tahan air IPX8 30m menawarkan perlindungan komprehensif untuk ponsel di laut dalam dengan tekanan tinggi<br><br>3.Bahan TPU anti kabut, tidak mudah menjadi kuning, anti sidik jari , memudahkan pengambilan gambar atau video dan pengenalan wajah.Telepon tetap respons sensitif bahkan di dalam tas tahan air<br><br>4.Bahan dari lanyard katun kuat, nyaman dipakai<br><br>5.Cocok untuk ponsel hingga 7,2 inci<br><br>Empat langkah mudah pemakaian:<br>1. Ikuti panah untuk memutar dan membuka penutup untuk mengeluarkan kantong anti air<br>2. Pisahkan kedua strip penutup tas, masukkan ponsel ke dalam tas, dan segel kembali strip penutup<br>3. Lipat strip segel dan dorong ke dalam silinder<br>4. Tutup penutup, siap digunakan<br><br>Berat Paket: 150 gram<br><br>SPEK:<br>Material: PC+TPU+cotton<br>Compatible: Handphone max 7.2 inches",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2021/7/8/cbbbcfa8-4e8c-4211-b8a9-5a06d608f80e.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2021/7/8/154e9faa-5c77-4779-9bdf-0fe27a7c6cc0.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2021/7/8/28f5943b-0cb4-4ff2-bd65-a9e72e016de9.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2021/7/8/cf5b05e6-3db6-4375-be9c-44694421abde.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2021/7/8/0284d2ea-cf87-4942-a0b1-a1ca07261b69.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "baseus-cylinder-slip-cover-waterproof-bag-phone-case-hp-anti-air-basah-hitam",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 250+",
    "country": "Kab. Tangerang",
    "estimation": "Estimasi tiba 31 Okt - 3 Nov",
    "toko": {
      "username": "baseonuserofficial",
      "toko_name": "Base On User Official Shop",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/5/29/cc124e49-1c5f-4c79-a07b-1a910be5e17a.jpg"
    }
  },
  {
    "product_name": "Softcase Anti Crack Elegant Casing Hp Samsung Galaxy A03 Free TG",
    "price": "Rp14.449",
    "stock": "1472",
    "description": "- Produk Tidak Semua Persis Seperti Gambar<br>- Barang yang anda bisa lihat ready stock dan bisa langsung dipesan.<br>- Jika barang yang dibeli ternyata habis, kami akan menggantikan dengan uang cash atau ditukar dengan barang lain yang harganya sama.<br>- Barang yang dijual lengkap sesuai dengan isi foto.<br>- Setiap orderan yang kami terima, kami usahakan proses paling lambat 2 X 24 jam.<br>- Tidak menerima Retur/Komplen barang yang rusak/cacat karena kami Sudah Mnegecek Brang sebelum Dikirim Dengan Kondisi Baik<br>- Untuk orderan yang masuk di hari Minggu akan kami proses di hari SENIN.<br>- CHECK dan PASTIKAN lebih teliti apakah alamat / nama / nomor handphone sudah sesuai, karena kami tidak menerima perubahan apapun jika sudah tercantum didalam invoice kami.<br>- Jika ada PERTANYAAN bisa ditanyakan di kolom diskusi atau pesan.<br>- Nomor Resi akan di input secepatnya.<br>Terima kasih telah mampir di toko kami, semoga bisa bekerja sama, salam sukses dan selamat berbelanja..",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/12/21/5eff1b86-4836-4974-a905-b7c9313e6add.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/12/21/68ccdbf4-df4a-4cf8-90e8-2025d6f80bb5.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/12/21/42bba1fa-f7ef-46ff-898d-1ce782c3e887.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/12/21/5eff1b86-4836-4974-a905-b7c9313e6add.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/12/21/5eff1b86-4836-4974-a905-b7c9313e6add.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/12/21/68ccdbf4-df4a-4cf8-90e8-2025d6f80bb5.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/12/21/42bba1fa-f7ef-46ff-898d-1ce782c3e887.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "softcase-anti-crack-elegant-casing-hp-samsung-galaxy-a03-free-tg",
    "discount": "4%",
    "discount_price": "Rp15.000",
    "rating": "5",
    "sold": "Terjual 28",
    "country": "Kota Bekasi",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "asmanofficial",
      "toko_name": "Asman official",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2019/9/29/6195551/6195551_9dc2fb26-14b4-4f11-905a-ffe8e262e71a.jpg"
    }
  },
  {
    "product_name": "Softcase Xiaomi Redmi S2 Shining Bear Casing Case hp - Hijau terang, Redmi S2",
    "price": "Rp25.000",
    "stock": "995",
    "description": "Selamat datang di toko kami ❤❤❤<br>Produk yang diperintahkan akan pabrik Indonesia dalam satu hari kerja, dan kami bergegas ke arahmu!<br><br>Kenapa kau memilih kami?<br>🤙KUALITAS TINGGI<br>Casing ponsel ini terbuat dari bahan Casing hp berkualitas tinggi serta, mudah dipasang &amp; dilepas, memberikan perlindungan terbaik untuk HP anda<br><br>🤙PENGIRIMAN LOKAL<br>Pengiriman langsung dari pabrik Indonesia untuk memastikan pengiriman barang yang cepat dan aman ke tangan Anda.<br><br>🤙LAYANAN YANG BAIK<br>Jika Anda memiliki pertanyaan tentang produk, layanan pelanggan akan meminta Anda online untuk layanan<br><br>Deskripsi Produk:<br>🤙💯Baru dan Dapat Dihapus.<br>🤙Desain pola kotak mode membuat ponsel Anda tetap pribadi dan terasa nyaman.<br>🤙Penyerap goncangan, bahan anti pecah.<br>🤙Lanyard menambahkan perlindungan terhadap jatuh yang tidak disengaja.<br>🤙Semua fungsi berfungsi dengan baik bahkan ponsel Anda dalam kasus ini.<br><br>Fitur:<br>Bahan: Terbuat dari Silikon Berkualitas Tinggi + TPU<br>Isi paket:<br>1 x Softcase elektroplating + 1 x Lanyard",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/7/14/a31c873c-c3e7-4da5-a5d6-e4c1dd749021.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/7/14/7cebcbc4-5364-48e0-b130-c5110e787de5.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/7/14/625dbbfe-8771-481e-ab0b-d8b857abb2ce.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/7/14/eb49bbb2-188f-4d38-9df2-3250eeddc8d0.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/7/14/ab773776-7bbd-4440-9089-86fa04427717.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "softcase-xiaomi-redmi-s2-shining-bear-casing-case-hp-hijau-terang-redmi-s2",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.9",
    "sold": "Terjual 30+",
    "country": "Kab. Tangerang",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "shopeper",
      "toko_name": "Trendi Case",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/4/29/1df532de-5202-4b3e-8db7-a42d37a779de.jpg"
    }
  },
  {
    "product_name": "Casing hp Samsung A04 A13 A23 A12 A22 A11 A50 A31 A32 A51 Kartu Kuromi - Style 1, Samsung A04",
    "price": "Rp24.000",
    "stock": "997",
    "description": "Product Description:<br><br>100% Brand new and removable.<br><br>Fashion design fits your phone completely.<br><br>High quality material makes your phone classic and elegant and feels good.<br><br>The whole pattern design makes your phone upscale.<br><br>Shock absorbing, soft silicone shatterproof material.<br><br>Unique monster lens design to avoid camera scratches.<br><br>Thickened camera hole adds protection against accidental drops.<br><br>Even if your phone is in this state, all functions will work fine.<br><br><br>Feature:<br>Material: Made of High Quality TPU<br><br>Package Content:<br>1 x Case",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/21/2d8e02b8-05c2-4674-9d66-f48acbea29b3.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/21/981680e5-8813-49f0-9429-f3f51d2b9337.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/21/1acbde51-ba0a-43a8-801c-bc68e1364792.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/21/585c2e25-5542-4637-b4a0-ac4f38bfbdea.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/21/be18ef4a-36c0-4d6b-9efe-edc99be076f7.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-samsung-a04-a13-a23-a12-a22-a11-a50-a31-a32-a51-kartu-kuromi-style-1-samsung-a04",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.8",
    "sold": "Terjual 70+",
    "country": "Kab. Tangerang",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "shopeper",
      "toko_name": "Trendi Case",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/4/29/1df532de-5202-4b3e-8db7-a42d37a779de.jpg"
    }
  },
  {
    "product_name": "Casing Hp Redmi Note 10 Pro Soft Case Plating Standing Astronot - Blue",
    "price": "Rp35.000",
    "stock": "980",
    "description": "Soft Case Import Casing Plating Dan Dudukan Handphone Astronot Premium<br><br>Ready Warna<br><br>- Black / Hitam<br>- Blue / Biru<br>- Purple / Ungu<br><br>Packing Menggunakan Kardus &amp; Bubble Warp<br><br>HIGH QUALITY AND PREMIUM DESIGN<br>- Case tidak hanya didesain dengan sangat elegant dan terlihat sangat mewah saat digunakan, namun case ini juga melindungi hp dengan baik dari benturan ataupun goresan<br>- Stand function (bisa diberdiriin)<br>- Menutupi sempurna bagian body samping dan belakang HP<br>- Mudah untuk dipasang dan dilepaskan<br>- Slim &amp; Fit<br>- Memberikan kesan mewah<br>- Bahan Berkualitas dan nyaman<br><br>Terima kasih..",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/16/d5cf332f-8fe3-4b28-9fe6-57e0f7c1ed39.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/16/dbac38fb-dbc1-44b5-b2d6-1d968712bdca.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/16/263c8aea-f030-4b9d-b35c-e3cc90c3faef.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/16/2d27d9a5-ef66-45ba-8125-424ed1c08b6a.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-redmi-note-10-pro-soft-case-plating-standing-astronot-blue-c1ffa",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 40+",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "fahryonemore",
      "toko_name": "FahryOneMore",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/9/7/3bb0f3ba-2d30-424a-885a-b4599e243647.jpg"
    }
  },
  {
    "product_name": "Casing Hp Tecno Spark 20 Pro+ Plus Soft Case Anti Licin Premium - Biru",
    "price": "Rp28.000",
    "stock": "999",
    "description": "NEW ARRIVAL<br>Premium Case Anti Fingerprint<br><br>Matte Case <br><br>Ready Colour:<br>* Black<br>* Blue Navy<br><br>Type: Carbon Case<br>Function: Dirt-resistant<br>Compatible Brand: Tecno<br>Design: Plain,Matte,Business,Vintage<br>Features: Real Matte Anti Fingerprint TPU Case<br><br>Spesifikasi &amp; Review Singkat :<br>* Produk Signature<br>* Desain Ultra tipis 0.6 mm dan perfect fit<br>* Desain minimalis dan maskulin<br>* Detail sangat tinggi dan sangat presisi<br>* Sangat ringan, lentur, tidak mudah kendor<br>* Sangat nyaman digenggam dalam waktu lama<br>* Lubang dan sisi tertutupi sangat rapi<br>* Melindungi benturan dari semua sisi<br>* Ciri khas dengan lis UV Edge Finish<br>* Warna tajam yang tahan lama (anti kusam)<br>* Material Solid TPU berkualitas terbaik<br>* Matte Finish yang halus sempurna<br>* Sisi ujung lebih keras untuk perlindungan lebih<br>* Tidak menganggu sinyal ponsel<br>* tidak membuat baret pada ponsel<br>* Coating Anti Jamur, Bebas jamur<br>* Coating Anti Minyak, Bebas bekas jari<br>* Coating Anti Kotor, Mudah Dibersihkan<br>* Coating Anti Gores (Sewajarnya)<br><br>Terima kasih.",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/3/cab90a30-433d-41ff-91c2-3a7bffab6b9f.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/3/f3fe3ae5-6f75-40f3-a205-bc6877a735a4.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/3/8c254ceb-acf2-4228-b84d-ffe616100c74.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/3/680c2ecf-d03a-4ee4-8e4b-d83adbdc3451.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-tecno-spark-20-pro-plus-soft-case-anti-licin-premium-biru-0d08a",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 6",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "fahryonemore",
      "toko_name": "FahryOneMore",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/9/7/3bb0f3ba-2d30-424a-885a-b4599e243647.jpg"
    }
  },
  {
    "product_name": "Casing Custom Case Anime Hololive Vtuber Aesthetic V1 [All Type HP] - No 1, Hardcase 2D",
    "price": "Rp41.888",
    "stock": "187",
    "description": "SELAMAT  DATANG DI TOKO KAMI...<br>Mohon dibaca ^_^<br><br>Cara Order :<br><br>1. Chat admin terlebih dahulu apakah stock tipe HP kalian tersedia<br>2. Konfirmasi desain yang akan anda pakai, bisa dari katalog, desain sendiri atau request<br>3. Kirim desain ke admin<br>4. Checkout<br>5. Bayar<br>6. Proses pengerjaan<br><br>----------------------------------------------------------------------------------<br><br>Jenis Bahan :<br><br>~ Hardcase 2D Glossy<br>Bahan Keras (Polycarbonate)<br>Cetakan gambar : Hanya dipermukaan saja<br>Bagian samping : Hitam<br>Body : Bagian Samping Saja<br><br>~ Softcase<br>Bahan Silikon (Lentur)<br>Cetakan gambar : Hanya dipermukaan saja<br>Bagian samping : Bisa Hitam Dan Bening<br>Body : Full<br><br>~ Luxury Glass Case<br>Bahan Fuze Glass<br>Cetakan gambar : Hanya dipermukaan saja<br>Bagian samping : Hitam<br>Bagian Permukaan : Kaca<br>Body : Full<br><br>Note : Luxury Glass Case Aurora hanya untuk type hp tertentu<br><br>----------------------------------------------------------------------------------<br><br>All Case :<br><br>1. Menggunakan tinta khusus sehingga gambar tidak mudah luntur atau pecah<br>2. Print Menggunakan mesin sinar UV dari Jepang ( Bukan Sticker )<br>3. Jaminan kualitas terbaik di kelasnya<br>4. Case memiliki kualitas terbaik dan lebih awet<br><br>----------------------------------------------------------------------------------<br><br>NB : <br>- Pengiriman mentahan desain melalui email agar gambar tidak pecah<br>- Proses pengerjaan 2-4 Hari<br>- Perlu diketahui Semua Produk di toko kami adalah barang Pre Order, bukan Ready Stock!! jadi barang akan di cetak setelah pemesanan.<br><br>----------------------------------------------------------------------------------<br><br>TERIMAKASIH SUDAH BERKUNJUNG DI TOKO KAMI DAN TERIMAKASIH ATAS KERJASAMANYA……..<br>HAPPY SHOPING…..",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/2/19/cd1275d3-5bbd-4714-b448-be2be6205759.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/8/18/be7308ad-7382-414a-9dd5-891995ebf623.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/8/18/b9da3410-489c-4b1b-9566-3aceeea6e895.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/8/18/3b0e2681-fbad-48ef-b130-361f397e547f.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/8/18/f3bd388b-86ea-4e44-9c8a-3254cebd093d.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-custom-case-anime-hololive-vtuber-aesthetic-v1-all-type-hp-no-1-hardcase-2d",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.9",
    "sold": "Terjual 40+",
    "country": "Kota Tangerang",
    "estimation": "Estimasi tiba 4 - 8 Nov",
    "toko": {
      "username": "wokecase",
      "toko_name": "WOKECASE",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2022/8/17/631154eb-8399-4df6-b893-4ec3c44125eb.jpg"
    }
  },
  {
    "product_name": "OPPO A31 A32 A33 A35 A38 A52 A53 A54 A55 A57 A58 4G 5G HOLOGRAM CASING VIRAL KESING HP CESING HANDPHONE SOFT HARD CASE - Hitam, OPPO A31",
    "price": "Rp15.000",
    "stock": "1.000",
    "description": "KELEBIHAN PRODUK :<br><br>Bahan material : Hybrid campuran softcase dan hardcase<br>- Bahan TPU Premium Bertekstur Lembut<br>- Nyaman di Genggam Tidak Licin<br>- Melindungi Dari Benturan Dengan Sempurna<br>- Anti Scratch (Tidak Mudah Tergores)<br>- Tidak Menganggu Sinyal Ponsel<br>- Tidak menimbulkan baret pada bodi ponsel<br>- Anti Fingerprint apabila jari kita berminyak atau kotor walaupun terkena Case akan sangat mudah untuk dibersihkan<br><br>MEMBELI = SETUJU<br>tidak menerima returan jika tidak ada video unboxing<br>#Liang's Store",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/b8d741d4-ab31-4070-ac7d-fb632b1342bf.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/f80991f7-d666-4123-b4d7-acfadf046614.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/2a937aa8-a0c7-48d9-b2ab-e51c8e514d85.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/27ce4735-bedc-43cd-8564-548bb2336c1f.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "oppo-a31-a32-a33-a35-a38-a52-a53-a54-a55-a57-a58-4g-5g-hologram-casing-viral-kesing-hp-cesing-handphone-soft-hard-case-hitam-oppo-a31-b6b45",
    "discount": "null",
    "discount_price": "null",
    "rating": "null",
    "sold": "null",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "liangsstore",
      "toko_name": "Liang's Store",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2022/10/11/39000dc2-26cd-4750-988d-968d4c5b8bed.jpg"
    }
  },
  {
    "product_name": "REDMI NOTE 12 12 PRO 13 13 PRO 4G 5G HOLOGRAM CASING VIRAL KESING HP CESING HANDPHONE SOFT HARD CASE XIAOMI - Hitam, RDMI NOTE 12 4G",
    "price": "Rp15.000",
    "stock": "1.000",
    "description": "KELEBIHAN PRODUK :<br><br>Bahan material : Hybrid campuran softcase dan hardcase<br>- Bahan TPU Premium Bertekstur Lembut<br>- Nyaman di Genggam Tidak Licin<br>- Melindungi Dari Benturan Dengan Sempurna<br>- Anti Scratch (Tidak Mudah Tergores)<br>- Tidak Menganggu Sinyal Ponsel<br>- Tidak menimbulkan baret pada bodi ponsel<br>- Anti Fingerprint apabila jari kita berminyak atau kotor walaupun terkena Case akan sangat mudah untuk dibersihkan<br><br>MEMBELI = SETUJU<br>tidak menerima returan jika tidak ada video unboxing<br>#Liang's Store",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/317b15af-f0f5-4061-ad07-45e8c0389436.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/8db81a2c-8082-4cee-b521-d225092551dd.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/8f92280d-c2e1-4a8b-ac2e-76e10dea2e22.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/dc49bab4-24b3-4dac-b4a6-1f97f70a4cc0.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "redmi-note-12-12-pro-13-13-pro-4g-5g-hologram-casing-viral-kesing-hp-cesing-handphone-soft-hard-case-xiaomi-hitam-rdmi-note-12-4g-33014",
    "discount": "null",
    "discount_price": "null",
    "rating": "null",
    "sold": "null",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "liangsstore",
      "toko_name": "Liang's Store",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2022/10/11/39000dc2-26cd-4750-988d-968d4c5b8bed.jpg"
    }
  },
  {
    "product_name": "OPPO A1K A3S A5 A5S A7 A9 A11K A12 A15 A15S A16 A17 A17K A18 2020 HOLOGRAM CASING VIRAL KESING HP CESING HANDPHONE SOFT HARD CASE - Hitam, OPPO A1K",
    "price": "Rp15.000",
    "stock": "1.000",
    "description": "KELEBIHAN PRODUK :<br><br>Bahan material : Hybrid campuran softcase dan hardcase<br>- Bahan TPU Premium Bertekstur Lembut<br>- Nyaman di Genggam Tidak Licin<br>- Melindungi Dari Benturan Dengan Sempurna<br>- Anti Scratch (Tidak Mudah Tergores)<br>- Tidak Menganggu Sinyal Ponsel<br>- Tidak menimbulkan baret pada bodi ponsel<br>- Anti Fingerprint apabila jari kita berminyak atau kotor walaupun terkena Case akan sangat mudah untuk dibersihkan<br><br>MEMBELI = SETUJU<br>tidak menerima returan jika tidak ada video unboxing<br>#Liang's Store",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/e8335d3f-690a-4e22-b09c-ddd32ebe3bc5.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/69f9642c-046c-41dc-ba48-543275831b0a.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/a6dd3b05-1093-4a08-8100-eac7bc3df1b4.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/31/32e3ecef-6134-4b6e-ab8f-2dc886ac8099.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "oppo-a1k-a3s-a5-a5s-a7-a9-a11k-a12-a15-a15s-a16-a17-a17k-a18-2020-hologram-casing-viral-kesing-hp-cesing-handphone-soft-hard-case-hitam-oppo-a1k-ef9b1",
    "discount": "null",
    "discount_price": "null",
    "rating": "null",
    "sold": "3 barang berhasil terjual",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "liangsstore",
      "toko_name": "Liang's Store",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2022/10/11/39000dc2-26cd-4750-988d-968d4c5b8bed.jpg"
    }
  },
  {
    "product_name": "Samsung A8 Plus softcase casing hp cover silikon ultra thin TPU Case - Black, A8 Plus",
    "price": "Rp12.000",
    "stock": "39",
    "description": "READY UNTUK :<br><br>- A8 Plus<br><br>Warna : Black<br><br><br>INFO PENTING WAJIB !!!<br><br>1. Setiap pembukaan packing WAJIB MELAKUKAN VIDEO UNBOXING PAKET, jika produk rusak atau jumlahnya kurang akan divalidasi melalui video, jika tidak ada bukti video maka dianggap TIDAK SAH untuk komplain<br><br>2. PERHATIAN : KAMI TIDAK MENERIMA KOMPLAIN untuk produk rusak / pecah pada saat diterima ataupun rusak pada saat pengiriman. Jadi mohon kerjasamanya untuk membeli KARDUS PACKING PENGAMAN TAMBAHAN untuk meminimalisir terjadinya kerusakan pada pengiriman.<br><br>HARAP BACA SEBELUM BELI, PENTING !!!<br>- Semua Pesanan akan dikirim SESUAI VARIASI yang dipilih pembeli<br>- Tidak menerima request variasi via chat, catatan, ataupun hal lainnya<br>- MEMBELI = SETUJU<br><br>KAPAN PESANAN DIKIRIM ?<br>1. SENIN - JUMAT Order yang masuk sebelum Jam 15.00 akan dikirim pada hari itu juga, KECUALI Sameday &amp; Instant maks. jam 14.00<br>2. SABTU maks Jam 11.00 akan dikirim pada hari itu juga<br>3. MINGGU dan TANGGAL MERAH LIBUR !!!<br><br>DESKRIPSI PRODUK :<br>- Bahan softcase TPU lentur<br>- Ketebalan 0.9mm sangat tipis<br>- Ultra Thin<br>- Tekstur matte / doff tidak licin<br>- Camera lens Protection (Proteksi Camera)<br>- Tidak membekas sidik jari (Anti Fingerprint)<br>- Tombol dan lubang-lubang presisi",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/17/bbc3e5cf-afc3-432f-9f4c-4d291bca686c.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/17/c4176bf0-49bd-4a08-83c3-3bf8741e9d54.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/17/afe35655-2069-49ad-8e33-3b061cd10bb5.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/17/faa0eb68-3d0e-49af-b310-9f16c43c12b2.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/4/17/7b70f6d9-d320-484d-9467-4ca4d0b742b6.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "samsung-a8-plus-softcase-casing-hp-cover-silikon-ultra-thin-tpu-case-black-a8-plus",
    "discount": "40%",
    "discount_price": "Rp20.000",
    "rating": "4.9",
    "sold": "Terjual 26",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "casepedias",
      "toko_name": "Case Pedia",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2017/1/5/1042576/1042576_4d8ee5d1-ccf0-4adc-aac9-5fbf9eeadc99.jpg"
    }
  },
  {
    "product_name": "Trohe Kantong Casing HP Anti Air Mobile Phone Pouch Waterproof CaseBag - White",
    "price": "Rp27.000",
    "stock": "500",
    "description": "【Sorotan yang tidak boleh dilewatkan】<br>1.0.3mm tipis dan ringan, operasi layar sentuh：Menggunakan bahan TPU film tipis dan permeabel tinggi kamera lebih jernih dan layar sentuh sensitif.<br>2.Tekstur linen ketahanan yang kuat：Lembut dan ramah kulit, cocok untuk menyelam, berenang, pemandian air panas.<br>3.Anti-tenggelam anti-hilang tidak ada air masuk：Desain multi-segel, disegel lapis demi lapis, tidak takut air di ponsel.<br>4.Tahan air IPX8：Sangat tahan air, bawa ponsel Anda dan nikmati waktu hiburan.<br>5.Jendela besar tidak menghalangi kamera：Lensa ultra-tipis dan tinggi hanya berjarak 0.3mm dari bare metal, merekam momen indah dari jarak dekat.<br><br>【Parameter produk】<br>Nama Produk: tas tahan air ponsel<br>Panjang tali: 45cm<br>Spesifikasi: 22.5(panjang) x11cm (lebar)<br>Sertifikasi tahan air: Lembaga SGS 30m IPX8<br>Tampilan layar: 17(panjang) x8.2cm (lebar)<br>Model yang cocok: universal di bawah 7.2 inci<br><br>【Selamat datang di toko elektronik TROHE】<br>TROHE adalah toko produk elektronik yang mengintegrasikan produksi dan desain.Kami berkomitmen untuk menyediakan produk dengan kualitas terbaik dan layanan yang paling memuaskan bagi konsumen. Berharap untuk menjadi mitra jangka panjang Anda!<br>Untuk merayakan masuknya Trohe secara resmi ke Indonesia, kami mengadakan acara pembukaan toko baru. Tidak boleh dilewatkan! Ingatlah untuk mengikuti kami!",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/11/11/0faaf8c2-44f2-4ebe-b3f6-3cecfbe3552a.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/10/31/68fecafe-fdce-45fe-be6c-60c653138a26.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/10/31/9f1209a3-0038-448f-b90d-9615a9716ee6.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/10/31/84dccfef-486e-42a3-857a-c961de5235a1.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "trohe-kantong-casing-hp-anti-air-mobile-phone-pouch-waterproof-casebag-white",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.9",
    "sold": "Terjual 5 rb+",
    "country": "Kab. Tangerang",
    "estimation": "Estimasi tiba 31 Okt - 3 Nov",
    "toko": {
      "username": "trohe",
      "toko_name": "TROHE",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2021/10/17/9f779d26-091c-4076-89b7-bb9175c6b787.jpg"
    }
  },
  {
    "product_name": "Case Samsung Galaxy M33 5G Slim Matte Black Casing Handphone",
    "price": "Rp17.000",
    "stock": "285",
    "description": "READY STOCK<br>Slim Case Matte Samsung M33 5G<br><br>Warna:<br>* HITAM<br><br>Deskripsi:<br>- Bahan karet lentur dan nyaman<br>- Nyaman di genggam tidak licin<br>- Sangat Tipis hanya 0.7 mm dan ringan<br>- Desain minimalis dan berkesan maskulin<br>- Melindungi semua sisi<br>- Lubang dan sisi tertutupi rapi<br>- Material Solid TPU berkualitas sangat lentur<br>- Tidak menganggu terhadap sinyal ponsel<br>- Tidak menimbulkan baret pada bodi ponsel<br><br>Terima kasih<br>FahryOneMore",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/4/27/3d4c3e59-d39e-45b4-88ee-ca1f91bca657.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/4/27/37f7b8f9-43ee-45d6-bce0-e756a37b3693.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/4/27/9c30a996-e71e-46f9-bbe5-38bd3aaefa67.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/4/27/16222efe-650e-4bc1-8a15-9e91926d06c9.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/4/27/ee9423ba-fc1b-47e0-8d97-88ea64439bc4.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "case-samsung-galaxy-m33-5g-slim-matte-black-casing-handphone",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 100+",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "fahryonemore",
      "toko_name": "FahryOneMore",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/9/7/3bb0f3ba-2d30-424a-885a-b4599e243647.jpg"
    }
  },
  {
    "product_name": "Casing Hp OPPO A5 2020 A9 2020 Soft Case Handphone Silicone Case - 040 - 1, A5 2020/A9 2020",
    "price": "Rp17.500",
    "stock": "232",
    "description": " TERSEDIA CASING CASE PREMIUM TIPE LENGKAP TERSEDIA UNTUK<br>#INFINIX<br>#OPPO<br>#REALME <br>#SAMSUNG<br>#VIVO<br>#IPHONE<br>#ADVAN<br>#XIAOMI<br>#ASUS<br>JIKA INGIN ORDER TIPE LAIN TULIS DI CATATAN YA KAK JADI KAKAK TINGGAL PILIH GAMBAR MANA YANG KAKAK SUKA . LALU UNTUK TIPE HP NYA DI CANTUMIN DI CATATAN ATAU BISA LANGSUNG CANTUMIN TIPE NYA DI CHAT YA.. CONTOH ORDER (NO 1 UNTUK OPPO Y12S)",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/10/15/4485a939-8e5d-4f64-a1a6-5dd8d05a2409.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/10/15/196db001-875a-401d-be29-220b43e844f1.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/10/15/dbbb36ff-9bd9-49e1-b9c0-2e2f1decc1ef.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/10/15/6fc36222-2677-4204-b1d1-ab2b68b10b53.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/10/15/407b8276-8c4c-4497-97fb-54556706eabc.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-oppo-a5-2020-a9-2020-soft-case-handphone-silicone-case-040-1-a5-2020-a9-2020",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.9",
    "sold": "Terjual 40+",
    "country": "Kota Bogor",
    "estimation": "Estimasi tiba besok - 2 Nov",
    "toko": {
      "username": "intristore",
      "toko_name": "intristore",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2021/9/8/3677650a-38dd-4c90-a675-9e651b3cd681.jpg"
    }
  },
  {
    "product_name": "Casing Hp Vivo Y12s Y20 Y20s Y20A Soft Case plating Dudukan Astronot - Blue, Vivo Y20",
    "price": "Rp35.000",
    "stock": "92",
    "description": "Soft Case Import Casing Plating Dan Dudukan Handphone Astronot Premium<br><br>Ready Warna<br><br>- Black / Hitam<br>- Blue / Biru<br>- Purple / Ungu<br><br>Packing Menggunakan Kardus &amp; Bubble Warp<br><br>High Quality Premium<br><br>- Case tidak hanya didesain dengan sangat elegant dan terlihat sangat mewah saat digunakan, namun case ini juga melindungi hp dengan baik dari benturan ataupun goresan<br>- Stand function (bisa diberdiriin)<br>- Menutupi sempurna bagian body samping dan belakang HP<br>- Mudah untuk dipasang dan dilepaskan<br>- Slim &amp; Fit<br>- Memberikan kesan mewah<br>- Bahan Berkualitas dan nyaman<br><br>Terima kasih..",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/7/12/7bbadef4-c987-4d0c-abd9-0fd657bca02c.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/7/12/e19dd6c8-905d-431e-812b-524252c17fdb.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/7/12/0be432e8-c3a9-420c-a1f6-78a1918d5da8.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/7/12/902c9e5c-8d6f-480a-9b8b-55405eb8b3fc.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-vivo-y12s-y20-y20s-y20a-soft-case-plating-dudukan-astronot-blue-vivo-y20-3aea0",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 70+",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "fahryonemore",
      "toko_name": "FahryOneMore",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/9/7/3bb0f3ba-2d30-424a-885a-b4599e243647.jpg"
    }
  },
  {
    "product_name": "Casing Hp Vivo Y22 Y22s Soft Case Plating Standing Astronot - Black, Vivo Y22",
    "price": "Rp35.000",
    "stock": "477",
    "description": "Soft Case Import Casing Plating Dan Dudukan Handphone Astronot Premium<br><br>Ready Warna<br><br>- Black / Hitam<br>- Blue / Biru<br>- Purple / Ungu<br><br>Packing Menggunakan Kardus &amp; Bubble Warp<br><br>HIGH QUALITY AND PREMIUM DESIGN<br>- Case tidak hanya didesain dengan sangat elegant dan terlihat sangat mewah saat digunakan, namun case ini juga melindungi hp dengan baik dari benturan ataupun goresan<br>- Stand function (bisa diberdiriin)<br>- Menutupi sempurna bagian body samping dan belakang HP<br>- Mudah untuk dipasang dan dilepaskan<br>- Slim &amp; Fit<br>- Memberikan kesan mewah<br>- Bahan Berkualitas dan nyaman<br><br>Terima kasih..",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/8/c8773fbf-364d-4ec9-8106-d07d9a8fb757.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/8/0cabe90c-185f-490d-b276-bb0eb56b509b.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/8/11aca38a-7ff3-47ac-80bb-57868865e6cc.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/8/8/4b96260e-682e-4b46-85c1-3a03e9dcf128.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-vivo-y22-y22s-soft-case-plating-standing-astronot-black-vivo-y22-57f4c",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 70+",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "fahryonemore",
      "toko_name": "FahryOneMore",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/9/7/3bb0f3ba-2d30-424a-885a-b4599e243647.jpg"
    }
  },
  {
    "product_name": "LOKMAT APPLLP MAX Jam Tangan Pintar Pria, 4G WiFi, Kamera Ganda 4/64",
    "price": "Rp3.399.000",
    "stock": "2",
    "description": "ALL&nbsp;OPERATOR&nbsp;4G&nbsp;SIM&nbsp;CARD<br><br>Bisa&nbsp;game&nbsp;playstore&nbsp;PUBG&nbsp;ML<br>Whatsapp&nbsp;Call&nbsp;Video&nbsp;Call&nbsp;Tanpa&nbsp;Kartu&nbsp;terhubung&nbsp;nomor&nbsp;smartphone<br>Bisa&nbsp;lepas&nbsp;pasang&nbsp;Strap&nbsp;jadi&nbsp;Android&nbsp;Phone<br>Notifikasi Pair Hp &amp; Watch ( Android )<br>Dapat&nbsp;Fitur&nbsp;kesehatan&nbsp;<br>Garansi&nbsp;1&nbsp;tahun<br><br>Note&nbsp;:&nbsp;Pertama&nbsp;kali&nbsp;masuk&nbsp;Play&nbsp;Store&nbsp;akan&nbsp;dilakukan&nbsp;update&nbsp;playstore,&nbsp;tunggu&nbsp;2&nbsp;menit&nbsp;dan&nbsp;Playstore&nbsp;siap&nbsp;digunakan.<br><br>OS:&nbsp;Android&nbsp;9.0<br>Memory:&nbsp;RAM&nbsp;4GB&nbsp;+&nbsp;ROM&nbsp;64GB<br>CPU:&nbsp;MTK6761<br>battery&nbsp;type:&nbsp;2300&nbsp;mAh&nbsp;Tahan&nbsp;3&nbsp;hari<br>Fitur:&nbsp;Play&nbsp;Store,&nbsp;3G&nbsp;4G,&nbsp;Akses&nbsp;Internet&nbsp;Nirkabel,&nbsp;Navigasi&nbsp;Google&nbsp;Map,&nbsp;&nbsp;GPS,&nbsp;Playback&nbsp;MP3,&nbsp;Layar&nbsp;Sentuh,&nbsp;Email,&nbsp;Radio&nbsp;FM,&nbsp;Auto&nbsp;Fokus,&nbsp;Built&nbsp;in&nbsp;flash.<br>Water&nbsp;Resist&nbsp;:&nbsp;IP67&nbsp;Tahan&nbsp;Air<br>Resolusi&nbsp;layar:&nbsp;480X640<br>display:&nbsp;2.88-inch&nbsp;TFT&nbsp;Full&nbsp;Touch&nbsp;Color&nbsp;Screen&nbsp;480*640&nbsp;<br>Kamera&nbsp;7MP&nbsp;primary,&nbsp;dual&nbsp;camera<br>Raise&nbsp;To&nbsp;Wake&nbsp;<br>Notifikasi&nbsp;Pair&nbsp;Hp&nbsp;&amp;&nbsp;Watch&nbsp;Bluetooth&nbsp;(android)<br>&nbsp;<br><br>Compatible&nbsp;system:Compatible&nbsp;with&nbsp;Android&nbsp;6.0/&nbsp;iOS&nbsp;11.0&nbsp;and&nbsp;above<br>Warna&nbsp;layar:warnaCertification:ECE<br>battery&nbsp;type:2300&nbsp;mAh<br>Bahan&nbsp;pita:RubberNama&nbsp;<br>merek:LokmatTipe&nbsp;Gerakan:ElektronikProduct&nbsp;name:LOKMAT&nbsp;APPLLP&nbsp;MAX&nbsp;4G&nbsp;Android&nbsp;Smart&nbsp;Watch<br>Kamera:&gt;&nbsp;7&nbsp;MP<br>Nomor&nbsp;model:APPLLP&nbsp;MAXDesain:penggeserBahan&nbsp;Sarung:AloiCamera:5.0MP+13MPTempat&nbsp;asal:Guangdong,&nbsp;ChinaJenis&nbsp;tampilan:Monitor&nbsp;LCD&nbsp;KomputerPerincian&nbsp;Pengemasan:Selling&nbsp;Units:&nbsp;Single&nbsp;item&nbsp;Single&nbsp;package&nbsp;size:&nbsp;10X8X7&nbsp;cm&nbsp;Single&nbsp;gross&nbsp;weight:&nbsp;0.3&nbsp;kg&nbsp;Package&nbsp;Type:&nbsp;color&nbsp;boxjual&nbsp;Unit:Single&nbsp;itemukuran&nbsp;paket&nbsp;tunggal:12X10.5X7.1&nbsp;cm<br>berat&nbsp;kotor:0.350&nbsp;Kg",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/12/19/c7a9c79a-d509-4800-8342-bffd6dd001bf.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/8/13/7783025a-48dd-4913-9061-e79fb28770c7.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/12/19/c55dc630-0109-4f51-b342-524318e45f31.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/12/19/edd5b6a3-0dc7-4dbf-b306-da58042f038f.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "lokmat-appllp-max-jam-tangan-pintar-pria-4g-wifi-kamera-ganda-4-64",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.9",
    "sold": "Terjual 14",
    "country": "Kab. Bandung Barat",
    "estimation": "Estimasi tiba 31 Okt - 3 Nov",
    "toko": {
      "username": "valoz",
      "toko_name": "valoz",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2022/11/24/4aaf6a18-06e5-4421-8807-c04ae4fee6bb.jpg"
    }
  },
  {
    "product_name": "Case Hp OPPO A54 A55 A57 2022 A77s A58 4G Casing Hybrid Plating Hologram Hardcase - A54, Hitam / random",
    "price": "Rp35.000",
    "stock": "100",
    "description": "Tersedia untuk :<br>- OPPO A54<br>- OPPO A55 4G<br>- OPPO A57 2022 / A77s<br>- OPPO A58 4G<br><br>UNTUK MODEL / WARNA YANG DI MINTA / PILIH JIKA KOSONG KAMI AKAN KONFIRMASI LEWAT CHAT<br>APABILA TIDAK RESPON DALAM WAKTU 2-3 JAM MAKA AKAN DI KIRIM RANDOM ,MAKA TIDAK BISA COMPLAIN<br><br>KOMPLAIN HARUS MENYERTAKAN VIDEO UNBOXING TANPA PUTUS2<br><br>MEMBELI = SETUJU<br><br>HAPPY SHOPPING<br>CASE IMD PLATE HOLOGRAM<br><br>NOTE =UNTUK WARNA YANG DI MINTA / PILIH JIKA KOSONG AKAN DI KIRIM RANDOM KECUALI ADA CATATAN HITAM KOSONG CANCEL<br><br>KELEBIHAN PRODUK :<br><br>Bahan material : Hybrid campuran softcase dan hardcase<br>- Bahan TPU Premium Bertekstur Lembut<br>- Nyaman di Genggam Tidak Licin<br>- Melindungi Dari Benturan Dengan Sempurna<br>- Anti Scratch (Tidak Mudah Tergores)<br>- Tidak Menganggu Sinyal Ponsel<br>- Tidak menimbulkan baret pada bodi ponsel<br>- Anti Fingerprint apabila jari kita berminyak atau kotor walaupun terkena Case akan sangat mudah untuk dibersihkan<br><br>Beli = Setuju",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/10/22/1c824cb0-87fb-43a4-902e-342fd467e797.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/10/22/61f7041b-2cdd-46e1-9db0-bdffde02697c.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "case-hp-oppo-a54-a55-a57-2022-a77s-a58-4g-casing-hybrid-plating-hologram-hardcase-a54-hitam-random-2a1c8",
    "discount": "null",
    "discount_price": "null",
    "rating": "null",
    "sold": "null",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 2 Nov",
    "toko": {
      "username": "aksesorishitam",
      "toko_name": "Aksesoris Hitam",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2018/5/31/3300071/3300071_f88a0c19-b417-48e9-bef2-61a9719f8ac5.jpg"
    }
  },
  {
    "product_name": "Custom Case Softcase Hardcase Silikon Casing Hp Astronot V2 - 0006, Hardcase 3D",
    "price": "Rp35.500",
    "stock": "998",
    "description": "<br>READY UNTUK TIPE HP :<br>• IPHONE<br>• SAMSUNG<br>• OPPO<br>• XIAOMI<br>• VIVO<br>• REALME<br>• ASUS<br>• INFINIX<br><br>JENIS CASE :<br>• ANTICRACK<br>• BLACK MATTE<br>• HARDCASE <br><br>CARA ORDER :<br>Silahkan chat admin untuk tanya stok casing, Jika stok tersedia lakukan pembayaran dan kirim gambar yang di inginkan ke admin..<br>Jika ingin pake gambar sendiri bisa di kirim ke admin<br><br>PERHATIAN :<br>• Untuk proses cetak gambar sampai dengan pengiriman paling lambat membutuhkan waktu 3 hari (jika orederan overload) tidak teirmasuk hari minggu dan tanggal merah<br>• Jika salah tipe HP dari anda sendiri tidak bisa complain<br>• Jika salah tipe HP dari kita diganti baru dengan gambar yang sama dan foto yang sama<br>• Hasil cetak tergantung gambar asli<br>• Hasil cetak jika ada perbedaan warna adalah WAJAR, tidak bisa sama PERSIS<br>• Order berarti setuju",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/9/6/eb26b144-5fee-4c76-a907-a3bb2a50e84c.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/9/6/554232ec-aafb-4c3f-8c8e-e2c125333e8d.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/9/6/1dc92387-6fc8-4017-b11f-a192654b1632.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/9/6/b92bf5b1-58e2-40d1-aa3b-98b0dbb99762.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/9/6/6ae04b74-0b2e-48f3-9e73-74c27537e7e7.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "custom-case-softcase-hardcase-silikon-casing-hp-astronot-v2-0006-hardcase-3d",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.3",
    "sold": "Terjual 100+",
    "country": "Kab. Bogor",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "sanjayacases",
      "toko_name": "sanjayacases",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2021/7/12/11303933/11303933_aa6588b7-73dc-485c-acce-675385b003d6.jpg"
    }
  },
  {
    "product_name": "Casing Hp Samsung M15 5G Soft Case plating Dudukan Astronot - Purple",
    "price": "Rp35.000",
    "stock": "998",
    "description": "Soft Case Import Casing Plating Dan Dudukan Handphone Astronot Premium<br><br>Ready Warna<br><br>- Black / Hitam<br>- Blue / Biru<br>- Purple / Ungu<br><br><br>High Quality Premium<br><br>- Case tidak hanya didesain dengan sangat elegant dan terlihat sangat mewah saat digunakan, namun case ini juga melindungi hp dengan baik dari benturan ataupun goresan<br>- Stand function (bisa diberdiriin)<br>- Menutupi sempurna bagian body samping dan belakang HP<br>- Mudah untuk dipasang dan dilepaskan<br>- Slim &amp; Fit<br>- Memberikan kesan mewah<br>- Bahan Berkualitas dan nyaman<br><br>Terima kasih..",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/17/1c48d1a5-12f4-4b4f-a941-064430dd4113.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/17/c79f539c-3abb-456e-a965-0d1e7fb3dfff.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/17/a7223ee2-c1f9-4747-9dcb-3192a0b3c076.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/7/17/8928acbc-256e-4333-89fb-17ce7dc32e47.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-samsung-m15-5g-soft-case-plating-dudukan-astronot-purple-c052a",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 6",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "fahryonemore",
      "toko_name": "FahryOneMore",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/9/7/3bb0f3ba-2d30-424a-885a-b4599e243647.jpg"
    }
  },
  {
    "product_name": "Casing Hp Infinix Zero 5G Soft Case Plating Standing Astronot Premium - Blue, Zero 5G 2023",
    "price": "Rp35.000",
    "stock": "999",
    "description": "Soft Case Import Casing Plating Dan Dudukan Handphone Astronot Premium<br><br>Ready Warna<br><br>- Black / Hitam<br>- Blue / Biru<br>- Purple / Ungu<br><br>Packing Menggunakan Kardus &amp; Bubble Warp<br><br>HIGH QUALITY AND PREMIUM DESIGN<br>- Case tidak hanya didesain dengan sangat elegant dan terlihat sangat mewah saat digunakan, namun case ini juga melindungi hp dengan baik dari benturan ataupun goresan<br>- Stand function (bisa diberdiriin)<br>- Menutupi sempurna bagian body samping dan belakang HP<br>- Mudah untuk dipasang dan dilepaskan<br>- Slim &amp; Fit<br>- Memberikan kesan mewah<br>- Bahan Berkualitas dan nyaman<br><br>Terima kasih..",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/6/6/6023c617-c2e3-4dcf-8cd4-925f8f11a819.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/6/6/fb4a7bca-1510-4443-98c7-79057bbb43a9.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/6/6/a123e160-2813-4398-ae78-f34d5d1c7bd4.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/6/6/ae53029a-003d-4527-ada9-022af2619c1a.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-infinix-zero-5g-soft-case-plating-standing-astronot-premium-blue-zero-5g-2023",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 6",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "fahryonemore",
      "toko_name": "FahryOneMore",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/9/7/3bb0f3ba-2d30-424a-885a-b4599e243647.jpg"
    }
  },
  {
    "product_name": "Case Casing hp iPhone 6 6S 6+ 6S+ 7 7+ 8 8+ Plus SE Berkelahi Luffy - Hijau terang, iPhone 6",
    "price": "Rp25.000",
    "stock": "999",
    "description": "Selamat datang di toko kami ❤❤❤<br>Produk yang diperintahkan akan pabrik Indonesia dalam satu hari kerja, dan kami bergegas ke arahmu!<br><br>Kenapa kau memilih kami?<br>🤙KUALITAS TINGGI<br>Casing ponsel ini terbuat dari bahan Casing hp berkualitas tinggi serta, mudah dipasang &amp; dilepas, memberikan perlindungan terbaik untuk HP anda<br><br>🤙PENGIRIMAN LOKAL<br>Pengiriman langsung dari pabrik Indonesia untuk memastikan pengiriman barang yang cepat dan aman ke tangan Anda.<br><br>🤙LAYANAN YANG BAIK<br>Jika Anda memiliki pertanyaan tentang produk, layanan pelanggan akan meminta Anda online untuk layanan<br><br>Deskripsi Produk:<br>🤙💯Baru dan Dapat Dihapus.<br>🤙Desain pola kotak mode membuat ponsel Anda tetap pribadi dan terasa nyaman.<br>🤙Penyerap goncangan, bahan anti pecah.<br>🤙Lanyard menambahkan perlindungan terhadap jatuh yang tidak disengaja.<br>🤙Semua fungsi berfungsi dengan baik bahkan ponsel Anda dalam kasus ini.<br><br>Fitur:<br>Bahan: Terbuat dari Silikon Berkualitas Tinggi + TPU<br>Isi paket:<br>1 x Softcase elektroplating + 1 x Tali Pengikat Silikon Lanyard",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/7/3/85bf40cf-a6f6-48aa-8749-0189f7428190.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/7/3/8020e5b1-94e6-414e-a950-092befc4ec54.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/7/3/4b976df0-d773-465c-a384-dba1e79c6581.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/7/3/de13edef-36dc-42a4-a91b-e63eb36db62c.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/7/3/7ce7ba96-1314-4dc3-bed0-296a06dcffb2.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "case-casing-hp-iphone-6-6s-6-6s-7-7-8-8-plus-se-berkelahi-luffy-hijau-terang-iphone-6",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 6",
    "country": "Kab. Tangerang",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "shopeper",
      "toko_name": "Trendi Case",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/4/29/1df532de-5202-4b3e-8db7-a42d37a779de.jpg"
    }
  },
  {
    "product_name": "BASEUS WATERPROOF CASING SARUNG PHONE HP HANDPHONE CASE ANTI AIR HUJAN - BLACK GREY",
    "price": "Rp99.000",
    "stock": "20",
    "description": "PERHATIAN : Ada 2 pilihan warna di iklan ini, dikirimkan sesuai dengan pilihan variant ya.. No komplain<br>-Tali Kuning (YELLOW GREY)<br>-Tali Hitam (BLACK GREY)<br><br>IPX8 Certified Waterproof, Responsive Touch Screen<br><br>Material : PVC + ABC + Silicone<br>Technology : Hot pressing<br>Applicable models : Phones under 7.2inch",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/8/23/e9b6c0d4-77d7-44fd-a271-da917697683a.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2020/10/9/22529bfb-e817-4762-8018-c5d8ffb34011.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2020/10/9/79b098c3-66d2-4d09-b3d9-315c662c4527.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2020/10/9/3b5e5439-99ea-4bf8-af32-513fbe8a0d4b.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2020/10/9/d264eb17-52da-442a-9706-4f62f31b030f.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "baseus-waterproof-casing-sarung-phone-hp-handphone-case-anti-air-hujan-black-grey",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.9",
    "sold": "Terjual 750+",
    "country": "Jakarta Utara",
    "estimation": "Estimasi tiba besok - 2 Nov",
    "toko": {
      "username": "koboitech",
      "toko_name": "KoboiTech",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2019/8/3/6514734/6514734_a37c1537-75ca-40f9-b18f-d165713df4fe.png"
    }
  },
  {
    "product_name": "Case Xiaomi 12 5G SoftCase Premium Casing Hp Slim Cover",
    "price": "Rp16.500",
    "stock": "8930",
    "description": "New Premium Ultra Slim Matte<br><br>Premium Soft Case<br><br>Spesifikasi produk :<br><br>- High Quality Premium Silicone<br>- Warna Doff<br>- Presisi 1000% Real Phone Model<br>- Terdapat List Di Bagian Samping<br>- Membuat Tampilan Menjadi Lebih Mewah dan Elegant<br>- Tombal empuk<br>- Tidak Licin Saat di Genggam ( Anti Fingerprint SYSTEM)<br><br>Available SoftCase Color:<br><br>- Black / Hitam Doft Matte <br><br>Standard Packaging kami sudah menggunakan kardus dan bubble wrap , jadi tidak perlu menambah biaya lagi agar pesanan aman sampai ke-tempat kamu.<br><br>Semua barang yang dikirim telah melalui proses Q.C tim kami guna menghindari pengiriman barang yang tidak sesuai Pesanan kamu.<br><br>Di Toko Kami senantiasa memenuhi kebutuhan untuk handphone kamu , XTRA BANYAK BONUS NYA , HARGA MIRING gak bikin KANTONG KERING, Kualitas DIJAMIN GAK GARING.",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/11/15/d510917c-b65f-4c53-9d78-940547a8d206.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/3/16/bf6d44c3-70be-4b1e-9802-ec2d7c6e536f.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/3/16/a635221b-2023-4bde-82c9-2c20acf0d5a8.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/3/16/ae2d69ad-ced1-4eba-95f0-6015b037bd97.png.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2022/3/16/27d402a9-cacf-494e-8ebe-eabced3c3524.png.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "case-xiaomi-12-5g-softcase-premium-casing-hp-slim-cover",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 70+",
    "country": "Jakarta Selatan",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "zonacaseimport",
      "toko_name": "Zona Case Import",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2017/7/22/2330759/2330759_de2c2461-8d56-4972-b1a9-a3e9feb9b09e.jpg"
    }
  },
  {
    "product_name": "Ringke iPhone 16 / 16 Plus / 16 Pro / 16 Pro Max Fusion Bold / Fusion Card / Fusion Magnetic Case Softcase Tipis Magsafe Anti Crack Casing Hardcase Slim Armor Military Drop Tahan Banting Sarung Protector HP - 16P-Mag-Clr",
    "price": "Rp0",
    "stock": "Habis",
    "description": "Official Store Ringke<br>✅Official Store ✅100% Original ✅Ready Stock ✅Buka setiap hari ✅Fast Delivery ✅Best Service<br><br>Mohon Perhatikan Varian Sebelum Pembelian<br>Tidak Menerima Order di Chat / Notes<br>Salah Pesan Tidak Bisa Retur<br><br>Varian Model:<br>&gt; iPhone 16 Pro Max<br>&gt; iPhone 16 Pro<br>&gt; iPhone 16 Plus<br>&gt; iPhone 16<br><br>Varian Casing:<br>✔ Fusion Bold<br>Anti Yellowing<br><br>✔ Fusion<br>Classic Design, Made in Korea<br><br>✔ Fusion Magnetic<br>Built-in Neodymium Magnet<br><br>✔ Fusion Card<br>Ada tempat untuk menaruh kartu setebal 2.2mm<br>Kompatibel dengan Wireless Charging (meskipun ada kartu di dalamnya)<br><br>Fitur:<br>▪️ Material: Hybrid, pinggir Softcase, belakang Hardcase<br>▪️ Raised Bezel, melindungi Layar dan Kamera dari gesekan<br>▪️ Support Wireless Charging<br>▪️ Melindungi dari debu, kotor, goresan<br>▪️ Semua sisi terlindungi (Full Coverage Protection)<br>▪️ Tactile button, tombol lebih clicky dan enak di tekan<br>▪️ Nyaman di genggam<br>▪️ Slim, Fashionable dan Minimalis<br>▪️ Anti-Cling Dot Matriks, No Watermark atau Efek Pelangi<br>▪️ Potongan Akurat dan Tombol Presisi<br>▪️ Tersedia lubang khusus untuk lanyard maupun neck strap<br>▪️ Loop Hole Lebih besar, dap<br>at menampung 2 buckle sekaligus<br>▪️ Kompatibel dengan beragam jenis screen protector",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/9/30/aa60d0ee-2485-4113-8873-8e376d3a96b0.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/9/30/202dc330-e3e5-43a3-9108-96b015d2dc56.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/9/30/cfe6a14b-522e-408f-b87b-6b2426900728.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/9/30/bf43ac32-e4ad-4923-916c-aa50ae0cf748.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/9/30/32213a9d-d115-4388-8bd6-ef76c6d94c17.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "ringke-iphone-16-16-plus-16-pro-16-pro-max-fusion-bold-fusion-card-fusion-magnetic-case-softcase-tipis-magsafe-anti-crack-casing-hardcase-slim-armor-military-drop-tahan-banting-sarung-protector-hp-16p-mag-clr-c881f",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 1",
    "country": "null",
    "estimation": "null",
    "toko": {
      "username": "ringke",
      "toko_name": "Official Ringke Partner",
      "image_toko": ""
    }
  },
  {
    "product_name": "anti crack infinix hot 10s softcase hp infinix hot 10T crack hot 10T",
    "price": "Rp4.000",
    "stock": "17",
    "description": "Ready All type Anti crack :<br>Samsung<br>Oppo<br>Vivo<br>Xiaomi<br>iPhone<br>Lenovo<br>Asus<br>Nokia<br>Huawei<br>Infinix<br> <br><br>Harap tanya sebelum order ya??<br>Terimakasih..<br>Happy shopping  🤗<br>",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2021/8/12/b326bd8a-1352-4b21-9d22-a9b09e4444d2.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2021/8/12/86e5a037-fd63-4ac8-a183-7eb52612427c.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "anti-crack-infinix-hot-10s-softcase-hp-infinix-hot-10t-crack-hot-10t",
    "discount": "null",
    "discount_price": "null",
    "rating": "5",
    "sold": "Terjual 60+",
    "country": "Jakarta Pusat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "iraina",
      "toko_name": "iraina_hong",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2021/1/23/4466415/4466415_887e0cea-b09a-44d0-9f3e-61d5765b2ec1.jpg"
    }
  },
  {
    "product_name": "Case Ultraslim Black Doff Matte Softcase Hp Xiaomi Redmi Note 5A Prime",
    "price": "Rp10.500",
    "stock": "1482",
    "description": "- Produk Tidak Semua Persis Seperti Gambar<br>- Barang yang anda bisa lihat ready stock dan bisa langsung dipesan.<br>- Jika barang yang dibeli ternyata habis, kami akan menggantikan dengan uang cash atau ditukar dengan barang lain yang harganya sama.<br>- Barang yang dijual lengkap sesuai dengan isi foto.<br>- Setiap orderan yang kami terima, kami usahakan proses paling lambat 2 X 24 jam.<br>- Tidak menerima Retur/Komplen barang yang rusak/cacat karena kami Sudah Mnegecek Brang sebelum Dikirim Dengan Kondisi Baik<br>- Untuk orderan yang masuk di hari Minggu akan kami proses di hari SENIN.<br>- CHECK dan PASTIKAN lebih teliti apakah alamat / nama / nomor handphone sudah sesuai, karena kami tidak menerima perubahan apapun jika sudah tercantum didalam invoice kami.<br>- Jika ada PERTANYAAN bisa ditanyakan di kolom diskusi atau pesan.<br>- Nomor Resi akan di input secepatnya.<br>Terima kasih telah mampir di toko kami, semoga bisa bekerja sama, salam sukses dan selamat berbelanja..",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/1/3/0657a00b-4850-4709-b03d-34a540b7955c.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/1/3/39cae881-4ab7-4599-b43a-f11bc4edd23e.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/1/3/3c8624ab-a833-4a16-b1a9-bf29e08e2c4e.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "case-ultraslim-black-doff-matte-softcase-hp-xiaomi-redmi-note-5a-prime",
    "discount": "5%",
    "discount_price": "Rp11.000",
    "rating": "4.9",
    "sold": "Terjual 80+",
    "country": "Kota Bekasi",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "asmanofficial",
      "toko_name": "Asman official",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2019/9/29/6195551/6195551_9dc2fb26-14b4-4f11-905a-ffe8e262e71a.jpg"
    }
  },
  {
    "product_name": "Softcase Ultraslim BlackMatte Casing Hp Xiaomi Redmi 5A - Free TG",
    "price": "Rp14.449",
    "stock": "1454",
    "description": "- Produk Tidak Semua Persis Seperti Gambar<br>- Barang yang anda bisa lihat ready stock dan bisa langsung dipesan.<br>- Jika barang yang dibeli ternyata habis, kami akan menggantikan dengan uang cash atau ditukar dengan barang lain yang harganya sama.<br>- Barang yang dijual lengkap sesuai dengan isi foto.<br>- Setiap orderan yang kami terima, kami usahakan proses paling lambat 2 X 24 jam.<br>- Tidak menerima Retur/Komplen barang yang rusak/cacat karena kami Sudah Mnegecek Brang sebelum Dikirim Dengan Kondisi Baik<br>- Untuk orderan yang masuk di hari Minggu akan kami proses di hari SENIN.<br>- CHECK dan PASTIKAN lebih teliti apakah alamat / nama / nomor handphone sudah sesuai, karena kami tidak menerima perubahan apapun jika sudah tercantum didalam invoice kami.<br>- Jika ada PERTANYAAN bisa ditanyakan di kolom diskusi atau pesan.<br>- Nomor Resi akan di input secepatnya.<br>Terima kasih telah mampir di toko kami, semoga bisa bekerja sama, salam sukses dan selamat berbelanja..",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/1/4/13f3238e-457f-482b-bb26-c2da627b4efe.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/1/4/301c9a17-48b0-4dfa-bce3-524e46a554ab.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/1/4/60f44270-cd9d-4bf4-af0f-34a7625271d8.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "softcase-ultraslim-blackmatte-casing-hp-xiaomi-redmi-5a-free-tg",
    "discount": "4%",
    "discount_price": "Rp15.000",
    "rating": "4.7",
    "sold": "Terjual 100+",
    "country": "Kota Bekasi",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "asmanofficial",
      "toko_name": "Asman official",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2019/9/29/6195551/6195551_9dc2fb26-14b4-4f11-905a-ffe8e262e71a.jpg"
    }
  },
  {
    "product_name": "Case Ringke Samsung Galaxy S22 / S22 Plus / S22 Ultra Fusion Onyx Fusion X Casing Softcase Anti Crack Tipis Slim Armor Military Drop Tahan Banting Pelindung Handphone Sarung HP TPU Cover Silikon Hardcase Clear Bening Transparan - S22U Onx-Navy",
    "price": "Rp119.850",
    "stock": "21",
    "description": "Official Store Ringke<br>✅Official Store ✅100% Original ✅Ready Stock ✅Buka setiap hari ✅Fast Delivery ✅Best Service<br><br>Mohon Perhatikan Varian Sebelum Pembelian<br>Tidak Menerima Order di Chat / Notes<br>Salah Pesan Tidak Bisa Retur<br><br>Varian Casing:<br>✔ Fusion X<br>Rugged Design Made in Korea<br><br>✔ Fusion<br>Design Classic Made in Korea<br><br>✔ Fusion Card<br>Ada tempat untuk menaruh kartu setebal 2.2mm<br>Kompatibel dengan wireless charging (meskipun ada kartu di dalamnya)<br><br>✔ Onyx<br>Material Softcase Jelly<br><br>✔ Air<br>Material Full Softcase TPU<br><br>Fitur :<br>▪️ Raised Bezel, melindungi Layar dan Kamera dari gesekan<br>▪️ Support Wireless Charging<br>▪️ Melindungi dari debu, kotor, goresan<br>▪️ Semua sisi terlindungi (Full Coverage Protection)<br>▪️ Tactile button, tombol lebih clicky dan enak di tekan<br>▪️ Nyaman di genggam<br>▪️ Slim, Fashionable dan Minimalis<br>▪️ Anti-Cling Dot Matriks, No Watermark atau Efek Pelangi<br>▪️ Potongan Akurat dan Tombol Presisi<br>▪️ Tersedia lubang khusus untuk lanyard maupun neck strap<br>▪️ Kompatibel dengan beragam jenis screen protector",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2024/5/31/fbdb5ec7-b2f6-4415-8a0b-2a094df209a9.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/5/28/af8e305d-ee31-47cd-bee8-7e45b17097c3.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/5/28/26c4f30e-9a64-4797-9aa1-f127e99db85f.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/5/28/3b7eb578-30c3-4297-ac67-fb5a1fdfe2cf.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2024/5/28/8fbde5a0-17d9-4344-acf7-c88f8a8eb058.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "case-ringke-samsung-galaxy-s22-s22-plus-s22-ultra-fusion-onyx-fusion-x-casing-softcase-anti-crack-tipis-slim-armor-military-drop-tahan-banting-pelindung-handphone-sarung-hp-tpu-cover-silikon-hardcase-clear-bening-transparan-s22u-onx-navy-eb9de",
    "discount": "56%",
    "discount_price": "Rp273.000",
    "rating": "5",
    "sold": "Terjual 4",
    "country": "Jakarta Utara",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "ringkesurabaya",
      "toko_name": "Ringke Surabaya",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/5/4/09f20dd6-037e-4de5-a8bd-9bd8de5eac69.jpg"
    }
  },
  {
    "product_name": "Case Samsung A23 4G 5G A33 A53 A73 5G casing hp cover tpu CALFSKIN - Grey, A73 5G",
    "price": "Rp19.000",
    "stock": "17",
    "description": "READY UNTUK :<br><br>- Samsung A12/M12<br>- Samsung A13 4G/5G/A04s<br>- Samsung A20/A30<br>- Samsung A22 4G/M22/M32<br>- Samsung A23 4G/5G<br>- Samsung A30s/A50/A50s<br>- Samsung A32 4G<br>- Samsung A33 5G<br>- Samsung A51<br>- Samsung A52/A52s 5G<br>- Samsung A53 5G<br>- Samsung A7 2018/A750<br>- Samsung A71<br>- Samsung A72 4G/5G<br>- Samsung A73 5G<br>- Samsung S7 Edge<br>- Samsung S7 Flat<br>- Samsung S8<br>- Samsung S8 Plus<br>- Samsung S10 Plus<br>- Samsung S20 Ultra<br>- Samsung S21 Ultra<br>- Samsung S22 Plus<br>- Samsung S22 Ultra<br>- Samsung Note 10<br>- Samsung Note 10 Plus<br>- Samsung Note 20<br>- Samsung Note 20 Ultra<br>- Samsung Note 8<br>- Samsung Note 9<br><br>Warna : Black, Brown, Green, Grey<br><br><br>INFO PENTING WAJIB !!!<br><br>1. Setiap pembukaan packing WAJIB MELAKUKAN VIDEO UNBOXING PAKET, jika produk rusak atau jumlahnya kurang akan divalidasi melalui video, jika tidak ada bukti video maka dianggap TIDAK SAH untuk komplain<br><br>2. PERHATIAN : KAMI TIDAK MENERIMA KOMPLAIN untuk produk rusak / pecah pada saat diterima ataupun rusak pada saat pengiriman. Jadi mohon kerjasamanya untuk membeli KARDUS PACKING PENGAMAN TAMBAHAN untuk meminimalisir terjadinya kerusakan pada pengiriman.<br><br>HARAP BACA SEBELUM BELI, PENTING !!!<br>- Semua Pesanan akan dikirim SESUAI VARIASI yang dipilih pembeli<br>- Tidak menerima request variasi via chat, catatan, ataupun hal lainnya<br>- MEMBELI = SETUJU<br><br>KAPAN PESANAN DIKIRIM ?<br>1. SENIN - JUMAT Order yang masuk sebelum Jam 15.00 akan dikirim pada hari itu juga, KECUALI Sameday &amp; Instant maks. jam 14.00<br>2. SABTU maks Jam 11.00 akan dikirim pada hari itu juga<br>3. MINGGU dan TANGGAL MERAH LIBUR !!!<br><br>DESKRIPSI PRODUK :<br>- Bahan Vintage Leather lembut dan nyaman + TPU Silikon<br>- Full Lens Protection, melindungi bagian camera dari goresan<br>- Anti-Drop Protection, melindungi dari benturan pada saat terjatuh<br>- Ada lubang tali untuk gantungan<br>- Mudah dipasang dan dilepaskan<br>- Lentur dan Flexible tidak merusak body HP",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2022/12/29/07fdae9a-d372-401c-9a05-ac0459c73874.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/6/30/58e33707-ef4d-41b5-a3fa-d73fc18e684c.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/6/30/d779a7d4-1848-47ea-a5c5-e338df72555e.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/6/30/c876b8e0-9a62-4245-96f9-03ec84820d3c.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "case-samsung-a23-4g-5g-a33-a53-a73-5g-casing-hp-cover-tpu-calfskin-grey-a73-5g",
    "discount": "53%",
    "discount_price": "Rp40.000",
    "rating": "4.9",
    "sold": "Terjual 100+",
    "country": "Jakarta Barat",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "casepedias",
      "toko_name": "Case Pedia",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/shops-1/2017/1/5/1042576/1042576_4d8ee5d1-ccf0-4adc-aac9-5fbf9eeadc99.jpg"
    }
  },
  {
    "product_name": "Casing hp Samsung S10 + S20 Ultra FE S21 S22 5G Soft Case Bunga Seni - Gaya 1, Samsung S10",
    "price": "Rp15.000",
    "stock": "999",
    "description": "Selamat datang di toko kami ❤❤❤<br>Produk yang diperintahkan akan pabrik Indonesia dalam satu hari kerja, dan kami bergegas ke arahmu!<br><br>Kenapa kau memilih kami?<br>🤙KUALITAS TINGGI<br>Casing ponsel ini terbuat dari bahan softcase berkualitas tinggi serta, mudah dipasang &amp; dilepas, memberikan perlindungan terbaik untuk HP anda<br><br>🤙PENGIRIMAN LOKAL<br>Pengiriman langsung dari pabrik Indonesia untuk memastikan pengiriman barang yang cepat dan aman ke tangan Anda.<br><br>🤙LAYANAN YANG BAIK<br>Jika Anda memiliki pertanyaan tentang produk, layanan pelanggan akan meminta Anda online untuk layanan<br><br>Deskripsi Produk:<br>🤙💯Brand baru dan dapat dihapus.<br>🤙Desain wavy mode menjaga ponselmu pribadi dan merasa baik.<br>🤙Membebaskan pola yang sama hadiah pemilik kembali.<br>🤙Pemilik belakang mendukung lokasi telepon 360 derajat ketika digunakan.<br>🤙Penyerap kejutan, materi silikon lembut yang tak bisa hancur.<br>🤙Desain lensa unik untuk menghindari goresan kamera.<br>🤙Lubang kamera yang tebal menambah perlindungan terhadap titik kecelakaan.<br>🤙Bahkan jika ponselmu berada di kondisi ini, semua fungsi akan bekerja dengan baik.<br><br>Fitur:<br>Bahan: Dibuat dari TPU Kualitas Tinggi<br>Isi paket:<br>1 x Kasing",
    "image_url": [
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/5/19/5b1ff941-6e69-42b1-80a2-f550dd4b5d47.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/5/19/03221176-fbe2-40b1-a413-2efcf07f51a2.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/5/19/09117362-aed5-4051-94fd-8a89cff6c802.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/5/19/3bc1f0c3-9eb6-49a0-9427-6bd2dd1649e2.jpg.webp?ect=4g",
      "https://images.tokopedia.net/img/cache/500-square/hDjmkQ/2023/5/19/14291010-d962-45da-acf3-c26fda3ee1b6.jpg.webp?ect=4g"
    ],
    "category": "casing-hp",
    "slug": "casing-hp-samsung-s10-s20-ultra-fe-s21-s22-5g-soft-case-bunga-seni-gaya-1-samsung-s10",
    "discount": "null",
    "discount_price": "null",
    "rating": "4.5",
    "sold": "Terjual 6",
    "country": "Kab. Tangerang",
    "estimation": "Estimasi tiba besok - 3 Nov",
    "toko": {
      "username": "shopeper",
      "toko_name": "Trendi Case",
      "image_toko": "https://images.tokopedia.net/img/cache/215-square/GAnVPX/2023/4/29/1df532de-5202-4b3e-8db7-a42d37a779de.jpg"
    }
  }
]
`
