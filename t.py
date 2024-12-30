# Baca konten file
file_path = "./internal/db/dummy/casing.go"

with open(file_path, "r", encoding="utf-8") as file:
    content = file.readlines()

# Cari dan beri tahu di baris mana karakter U+2060 ditemukan
for line_number, line in enumerate(content, start=1):
    if "\u2060" in line:
        print(f"Karakter U+2060 ditemukan di baris {line_number}: {line.strip()}")
