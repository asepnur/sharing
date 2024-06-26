---
theme: default 
_class: lead
paginate: true
backgroundColor: #fff
backgroundImage: url('https://marp.app/assets/hero-background.svg')
header: "TDD & BDD"
---

# Tema 
Fundamental TDD dan BDD: Pengembangan Berbasis Tes dan Perilaku

--- 

# Lingkup Materi
1. **Pengantar TDD**
   - Apa itu TDD?
   - Tujuan dari TDD
   - TDD vs. Pengujian Tradisional

2. **Prinsip-prinsip TDD**
   - Siklus Red-Green-Refactor
   - Menulis Tes yang Efektif

---

3. **Manfaat TDD**
   - Kualitas Kode
   - Dokumentasi
   - Peningkatan Desain

4. **Contoh Kecil TDD (Golang)**
   - Contoh TDD 

---

5. **Konsep TDD Lanjutan**
   - Mocking dan Stubbing
   - Pengujian Integrasi

6. **Kesalahan dan Tips untuk Menguasai TDD**
   - Kesalahan Umum dan Cara Menghindarinya
   - Tips untuk TDD yang Efektif

---
7. **BDD**
   - Definisi BDD
   - Gherkin Language
   - Contoh BDD
---
8. **TDD x BDD**
   - Perbedaan
   - Apakah bisa dikombinasikan?
   - Contoh Real usecase

---
# 1. Pengantar TDD

## Apa itu TDD?
- **Definisi**: 
	- TDD adalah pendekatan software development di mana tes ditulis sebelum menulis program/kode.
	- Jadi, kita pastikan program/kode kita sesuai dengan yang diharapkan dari tes.
	- TDD menjadikan testing sebagain prinsip untuk membuat sebuah program atau kode.

---

- **Tujuan dari TDD**
	- Memastikan semua kode yang ditulis lulus tes dan sesuai ekspektasi.
	- Bugs atau kesalahan logika bisa ditemukan lebih awal.
	- Membantu membentuk budaya refactor dan melakukannya menjadi lebih percaya diri.
	- Kode tidak rentan terhadap error. 
	- Mengarahkan ke kode atau struktur kode yang lebih baik.
	- Menambah produktivitas (di awal akan lebih memakan waktu).

---

## TDD vs. Tes Tradisional
- **Pengujian Tradisional**: Kode dulu, tes kemudian.
- **TDD**
	- Tes dulu, baru kode. Ini bikin kode kita selalu teruji dan sesuai dengan kebutuhan.
	- Bukan hanya fokus pada tes, tapi juga pada aspek lain pada development proses (design pattern, koloborasi antar tim, dll)

---

# 2. Prinsip-prinsip TDD

### Siklus Red-Green-Refactor
1. **Red**: Tulis tes yang gagal, Menulis tes pertama.
2. **Green**: Tulis kode seminimal mungkin buat lulus tes.
3. **Refactor**: Perbaiki kode tanpa ubah perilaku.
  - **Referensi**: Martin, R.C. (2008). *Clean Code: A Handbook of Agile Software Craftsmanship*. Prentice Hall.

---

### Menulis Tes yang Efektif
- **Sederhana dan Jelas**: Tes harus mudah dibaca dan dipahami.
- **Terisolasi**: Setiap tes fokus pada satu perilaku atau fungsionalitas.
- **Dapat Diulang**: Tes harus selalu menghasilkan hasil yang sama.
- **Hindari Pengulangan**: Pengulangan tes yang tidak perlu akan menambah kerumitan tes.
  - **Referensi**: Freeman, S., & Pryce, N. (2009). *Growing Object-Oriented Software, Guided by Tests*. Addison-Wesley.

---

# 3. Manfaat TDD

### Kualitas Kode
- Kode jadi lebih teruji.
- Mengurangi bug dan kesalahan.
- Lebih percaya diri untuk melakukan perubahan kode.
  - **Referensi**: Erdogmus, H., Morisio, M., & Torchiano, M. (2005). "On the effectiveness of the test-first approach to programming". IEEE Transactions on Software Engineering, 31(3), 226-237.

---

### Dokumentasi
- Tes jadi dokumentasi buat kode.
- Memudahkan pengembang baru buat paham kode.

---

### Peningkatan Desain
- Mendorong praktik desain yang lebih baik.
- Membantu identifikasi dan perbaikan masalah desain lebih awal.
  - **Referensi**: Beck, K. (2002). *Test-Driven Development by Example*. Addison-Wesley.

---

# 4. Contoh Kecil TDD (Golang)

### Contoh TDD 
- Spek: Buatkan fungsi untuk memvalidasi Authentikasi
- Kriteria:
   - Fungsi menerima username dan password
   - Hanya akan berjalan jika keduanya ada
- [Code Base](https://github.com/asepnur/sharing/tree/main/example/testing/tdd).
---

# 5. Konsep TDD Lanjutan
### Mocking dan Stubbing
- Mocks: Buat simulasi perilaku objek nyata.
- Stubs: Memberikan respons yang sudah ditentukan buat panggilan fungsi.
  - **Referensi**: Freeman, S., & Pryce, N. (2009). Growing Object-Oriented Software, Guided by Tests. Addison-Wesley.

---

## Pengujian Integrasi
- Uji interaksi antara berbagai bagian sistem.
- Pastikan sistem bekerja secara keseluruhan.
- Lebih baik dikombinasikan dengan BDD.

---

# 6. Kesalahan dan Tips untuk Menguasai TDD 
## Kesalahan Umum dan Cara Menghindarinya
- Nulis tes yang terlalu kompleks.
- Mengabaikan refactoring setelah tes lulus.
   - **Referensi**: Martin, R.C. (2008). Clean Code: A Handbook of Agile Software Craftsmanship. Prentice Hall.

---

## Tips untuk TDD yang Efektif
- Mulai dengan tes kecil dan sederhana.
- Tingkatkan kompleksitas secara bertahap.
- Terus refactor kode.
---
# 7. BDD: Pengembangan Berbasis Perilaku
## Apa Itu BDD?
Behavior-Driven Development (BDD) adalah proses pengembangan perangkat lunak yang berfokus pada sudut pandang pengguna dengan menggunakan bahasa yang mudah dipahami orang awam untuk mendorong kolaborasi antara pengembang, QA, dan pemangku kepentingan non-teknis.
- Meningkatkan komunikasi dan pemahaman.
- Berfokus pada perilaku yang diharapkan dari perangkat lunak.
- Diperkenalkan oleh Dan North.

---

## Konsep Utama BDD
- **Kolaborasi**: Melibatkan semua pemangku kepentingan
- **Bahasa Umum**: Menggunakan Gherkin untuk menulis tes
- **Spesifikasi yang Dapat Dieksekusi**: Tes otomatis dan dapat dijalankan
- **Dokumentasi**: Tes berfungsi sebagai dokumentasi yang hidup

---

## Bahasa Gherkin
Gherkin adalah bahasa spesifik domain yang dapat dibaca oleh bisnis.
- Digunakan untuk menulis cerita pengguna dan skenario
- Mengikuti format sederhana

---

## Sintaks Gherkin
- **Fitur**: Menjelaskan fitur yang diuji
- **Skenario**: Menjelaskan kasus uji spesifik
- **Given**: Menetapkan konteks awal
- **When**: Menjelaskan tindakan atau peristiwa
- **Then**: Menjelaskan hasil yang diharapkan

---

### Contoh:
```gherkin
Feature: Registrasi Pengguna

Scenario: Registrasi berhasil dengan Email/Kata Sandi
  Given seorang pengguna dengan nama pengguna "testuser" dan kata sandi "password"
  When pengguna mendaftar dengan email dan kata sandi
  Then registrasi harus berhasil
```
---
# 8. **TDD x BDD**
## Perbedaan
## Apakah bisa dikombinasikan?
## Contoh Real usecase
- [Code Base](https://github.com/asepnur/sharing/tree/main/example/testing/combine).

---
# Terima kasih