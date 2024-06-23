---
marp: true
paginate: true
theme: default
header: "TDD dan BDD"
backgroundImage: url('https://marp.app/assets/hero-background.svg')
---

# Apa Itu BDD?
Behavior-Driven Development (BDD) adalah proses pengembangan perangkat lunak yang berfokus pada sudut pandang pengguna dengan menggunakan bahasa yang mudah dipahami orang awam untuk mendorong kolaborasi antara pengembang, QA, dan pemangku kepentingan non-teknis.
- Meningkatkan komunikasi dan pemahaman.
- Berfokus pada perilaku yang diharapkan dari perangkat lunak.
- Diperkenalkan oleh Dan North.

---

# Konsep Utama BDD
- **Kolaborasi**: Melibatkan semua pemangku kepentingan
- **Bahasa Umum**: Menggunakan Gherkin untuk menulis tes
- **Spesifikasi yang Dapat Dieksekusi**: Tes otomatis dan dapat dijalankan
- **Dokumentasi**: Tes berfungsi sebagai dokumentasi yang hidup

---

# Bahasa Gherkin
Gherkin adalah bahasa spesifik domain yang dapat dibaca oleh bisnis.
- Digunakan untuk menulis cerita pengguna dan skenario
- Mengikuti format sederhana

---

# Sintaks Gherkin
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
