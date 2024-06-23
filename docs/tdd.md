# 1. Pengantar TDD

### Apa itu TDD?
- **Definisi**: 
	- TDD adalah pendekatan software development di mana tes ditulis sebelum menulis program/kode.
	- Jadi, kita pastikan program/kode kita sesuai dengan yang diharapkan dari tes.
	- TDD menulis testing di awal.
- **Tujuan dari TDD**
	- Memastikan semua kode yang ditulis lulus tes dan sesuai ekspektasi.
	- Bugs atau kesalahan logika bisa ditemukan lebih awal.
	- Membantu membentuk budaya refactor dan melakukannya menjadi lebih percaya diri.
	- Tidak prone to error.
	- Mengarhkan ke kode atau struktur kode yang lebih baik.
	- Menambah produktivitas (di awal akan lebih memakan waktu).

### TDD vs. Tes Tradisional
- **Pengujian Tradisional**: Kode dulu, tes kemudian.
- **TDD**
	- Tes dulu, baru kode. Ini bikin kode kita selalu teruji dan sesuai dengan kebutuhan.
	- Bukan hanya fokus pada tes, tapi juga pada aspek lain pada development proses.

---

# 2. Prinsip-prinsip TDD

### Siklus Red-Green-Refactor
1. **Red**: Tulis tes yang gagal, Menulis tes.
2. **Green**: Tulis kode seminimal mungkin buat lulus tes.
3. **Refactor**: Perbaiki kode tanpa ubah perilaku.
  - **Referensi**: Martin, R.C. (2008). *Clean Code: A Handbook of Agile Software Craftsmanship*. Prentice Hall.

### Menulis Tes yang Efektif
- **Sederhana dan Jelas**: Tes harus mudah dibaca dan dipahami.
- **Terisolasi**: Setiap tes fokus pada satu perilaku atau fungsionalitas.
- **Dapat Diulang**: Tes harus selalu menghasilkan hasil yang sama.
  - **Referensi**: Freeman, S., & Pryce, N. (2009). *Growing Object-Oriented Software, Guided by Tests*. Addison-Wesley.
