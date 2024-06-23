---
# Tech Sharing: Pengembangan Berbasis Tes (TDD) Lanjutan
---

## Pengantar TDD
- Apa itu TDD?
- Sejarah dan Evolusi
- TDD vs. Pengujian Tradisional

---

## Prinsip-prinsip TDD
- Siklus Red-Green-Refactor
- Menulis Tes yang Efektif

---

## Manfaat TDD
- Kualitas Kode
- Dokumentasi
- Peningkatan Desain

---

## Alur Kerja TDD dalam Go (Golang)
- Menyiapkan Proyek Go untuk TDD
- Menulis Tes Pertama
- Mengimplementasikan Kode
- Refactoring

---

## Konsep TDD Lanjutan
- Mocking dan Stubbing
- Menguji Kode Asinkron
- Pengujian Integrasi

---

## Studi Kasus dan Praktik Terbaik
- Contoh Dunia Nyata
- Kesalahan Umum dan Cara Menghindarinya
- Tips untuk TDD yang Efektif

---

## 1. Pengantar TDD

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
    - Mengarahkan ke kode atau struktur kode yang lebih baik.
    - Menambah produktivitas (di awal akan lebih memakan waktu).

---

### TDD vs. Tes Tradisional
- **Pengujian Tradisional**: Kode dulu, tes kemudian.
- **TDD**
    - Tes dulu, baru kode. Ini bikin kode kita selalu teruji dan sesuai dengan kebutuhan.
    - Bukan hanya fokus pada tes, tapi juga pada aspek lain pada development proses.

---

## 2. Prinsip-prinsip TDD

### Siklus Red-Green-Refactor
1. **Red**: Tulis tes yang gagal, Menulis tes.
2. **Green**: Tulis kode seminimal mungkin buat lulus tes.
3. **Refactor**: Perbaiki kode tanpa ubah perilaku.
    - **Referensi**: Martin, R.C. (2008). Clean Code: A Handbook of Agile Software Craftsmanship. Prentice Hall.

---

### Menulis Tes yang Efektif
- **Sederhana dan Jelas**: Tes harus mudah dibaca dan dipahami.
- **Terisolasi**: Setiap tes fokus pada satu perilaku atau fungsionalitas.
- **Dapat Diulang**: Tes harus selalu menghasilkan hasil yang sama.
    - **Referensi**: Freeman, S., & Pryce, N. (2009). Growing Object-Oriented Software, Guided by Tests. Addison-Wesley.

---

## 3. Manfaat TDD

### Kualitas Kode
- Kode jadi lebih teruji.
- Mengurangi bug dan kesalahan.
    - **Referensi**: Erdogmus, H., Morisio, M., & Torchiano, M. (2005). "On the effectiveness of the test-first approach to programming". IEEE Transactions on Software Engineering, 31(3), 226-237.

---

### Dokumentasi
- Tes jadi dokumentasi buat kode.
- Memudahkan pengembang baru buat paham kode.

---

### Peningkatan Desain
- Mendorong praktik desain yang lebih baik.
- Membantu identifikasi dan perbaikan masalah desain lebih awal.
    - **Referensi**: Beck, K. (2002). Test-Driven Development by Example. Addison-Wesley.