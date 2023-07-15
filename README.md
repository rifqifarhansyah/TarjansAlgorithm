# Algoritma Tarjan untuk Pencarian SCC dan Bridge
<h2 align="center">
   <a>Implementasi Algoritma Tarjan</a>
</h2>
<hr>

## Table of Contents
1. [General Info](#general-information)
2. [Creator Info](#creator-information)
3. [Features](#features)
4. [Technologies Used](#technologies-used)
5. [Setup](#setup)
6. [Usage](#usage)
7. [Video Capture](#videocapture)
8. [Screenshots](#screenshots)
9. [Structure](#structure)
10. [Project Status](#project-status)
11. [Room for Improvement](#room-for-improvement)
12. [Acknowledgements](#acknowledgements)
13. [Contact](#contact)

<a name="general-information"></a>

## General Information
Sebuah aplikasi berbasis website sederhana yang dapat menerima masukan sebuah `file txt` atau `input text` dari sebuah gambaran suatu graf berarah untuk kemudian ditentukan `Strongly Connected Components (SCC)` dan `Bridges` dari graf tersebut dengan `Algoritma Tarjan`. Algoritma Tarjan diimplementasikan bersamaan dengan algoritma penelusuran `Depth First Search (DFS)`. Program ini dibuat untuk memenuhi tugas 6 seleksi Lab IRK tahun 2023.

### Tarjans Algorithm
`ALgoritma Tarjan` merupakan algoritma yang digunakan untuk mendeteksi `Strongly Connected Components (SCC)` dari suatu graf berarah. Algoritma ini berjalan dengan menggunakan algoritma penelusuran `Depth-First Search (DFS)` dan mencatat informasi penting tentang setiap simpul yang dikunjungi, seperti `indeks` dan `low link value`. 

### Complexity
`Kompleksitas waktu` Algoritma Tarjan adalah O(V+E), dimana V adalah jumlah simpul (vertices) dan E adalah jumlah sisi (edges) dalam graf. Algoritma ini cukup efisien dalam mengidentifikasi SCC dalam skala graf yang cukup besar.

### Modification
`Modifikasi yang dilakukan terhadap algoritma tarjan` untuk mendeteksi SCC dari suatu graf melibatkan penambahan struktur data dan logika yang sesuai. Penulis menggunakan implementasi struktur data `GraphBridge` untuk mencapai tujuan tersebut.

Struktur data `GraphBridge` terdiri atas 2 komponen utama:`Nodes` dan `Bridges`. Komponen Nodes adalah sebuah map yang menyimpan simpul-simpul dalam graf beserta informasi yang terkait. Setiap simpul (`NodeBridge`) memiliki properti seperti `nama, indeks, low link value, status kunjungan (visited), dan simpul-simpul tetangga yang terhubung langsung (adjacent)`. Komponen `Bridges` adalah sebuah irisan yang digunakan untuk menyimpan `bridges` yang ditemukan.

Ketika algoritma Tarjan berjalan, penulis memanfaatkan struktur data GraphBridge untuk melacak dan mengidentifikasi bridges dalam graf. Pada setiap tahapan DFS, ketika suatu simpul dikunjungi, penulis memeriksa apakah simpul tetangga tersebut bukanlah `simpul parent` dari simpul saat ini. Jika simpul tetangga belum pernah dikunjungi sebelumnya, maka algoritma akan melakukan rekursi pada simpul tetangga tersebut. Selama proses ini, nilai `indeks` dan `low link value` dari simpul saat ini dan simpul tetangga akan diperbarui.

Jika nilai `low link value` dari simpul tetangga lebih besar dari indeks simpul saat ini, maka simpul tetangga tersebut merupakan ujung dari suatu `bridge`. penulis menyimpan informasi ini dalam struktur data `Bridge` yang kemudian ditambahkan ke dalam komponen `Bridges` dari `GraphBridge`.

Dengan adanya modifikasi ini, penulis dapat mengidentifikasi dan menyimpan `bridges` dalam graf menggunakan algoritma Tarjan. Informasi ini kemudian dapat digunakan untuk visualisasi atau analisis lebih lanjut.

Modifikasi yang penulis lakukan memungkinkan algoritma Tarjan untuk memiliki kemampuan tambahan dalam mendeteksi jembatan yang kuat, sehingga memberikan nilai tambah pada fungsionalitas algoritma tersebut.

### Edges
Ada beberapa jenis edges yang mungkin terdapat dalam suatu graf, yaitu:
1. Back Edge: Edge yang menghubungkan simpul ke simpul lain di atasnya dalam DFS tree.
2. Cross Edge: Edge yang menghubungkan simpul ke simpul lain di cabang yang berbeda dalam DFS tree.
3. Forward Edge: Edge yang menghbungkan simpul ke simpul lain di bawahnya dalam DFS tree.
4. Tree Edge: Edge yang menghubungkan simpul ke simpul anaknya dalam DFS tree.
 
<a name="creator-information"></a>

## Creator Information

| Nama                        | NIM      | E-Mail                      |
| --------------------------- | -------- | --------------------------- |
| Mohammad Rifqi Farhansyah   | 13521166 | 13521166@std.stei.itb.ac.id |

<a name="features"></a>

## Features
- Memilih `metode input` yang akan digunakan, yaitu: `file txt` atau `input text`
- Melakukan proses `pencarian SCC dan Bridge` dari masukan graf dengan `algoritma tarjan`
- Menampilkan hasil `visualisasi graf` dari hasil `masukan, SCC, dan Bridge`
- Menampilkan `Runtime Program` mulai dari pengguna menekan tombol `upload`, hingga diperoleh hasil `visualisasi`

<a name="technologies-used"></a>

## Technologies Used
* Backend: Golang
* Frontend: React
* Library yang digunakan:
    -  net/http: Digunakan untuk meng-handle permintaan HTTP pada server web.
    - os: Digunakan untuk mengoperasikan file dan direktori
    - bufio: Digunakan untuk membaca file baris per baris.
    - os/exec: Digunakan dalam menjalankan perintah shell untuk menghasilkan visualisasi graf menggunakan Graphviz.
    - axios: Digunakan untuk melakukan permintaan HTTP dari sisi klien.
    - @mui/material: Digunakan untuk melakukan styling pada kode frontend.

> Note:  You can use the latest version of the libraries.

<a name="setup"></a>

## Setup
1. Clone Repository ini dengan menggunakan command berikut
   ```sh
   https://github.com/rifqifarhansyah/TarjansAlgorithm.git
   ```
2. Buka Folder "TarjansAlgorithm" di Terminal
3. Install seluruh Packages yang diperlukan
4. Masuk ke folder `frontend` lalu jalankan dengan command via terminal
   ```sh
   npm start
   ```
5. Kemudian buka terminal kedua, masuklah ke folder `backend` dan jalankan command
   ```sh
   go run main.go
   ```
6. Buka `localhost` yang digunakan pada Browser Anda
7. Dalam kasus ini, `frontend=http://localhost:3000` dan `backend=http://localhost:8080`

<a name="usage"></a>

## Usage
1. Pilih `metode input` yang akan digunakan, yaitu: `file txt` atau `text input`
2. Tekan tombol `upload` untuk menjalankan pemrosesan program
3. `Graf` hasil pembacaan masukan akan otomatis ditampilkan di sisi kanan tampilan web
4. `Bridge` dan `SCC` akan otomatis ditampilkan di bagian bawah tombol `upload`
5. Tekan tombol `View Bridge` atau `View SCC` untuk menampilkan visualisasi-nya
6. `Visualisasi Bridge atau SCC` akan ditampilkan dengan mengoverwrite posisi `visualisasi graf awal`
7. `Runtime Program` akan ditampilkan pada bagian bawah web

<a name="videocapture"></a>

## Video Capture
<nl>

![TarjansAlgorithm Gif](https://github.com/rifqifarhansyah/TarjansAlgorithm/blob/main/img/TarjansAlgorithm.gif?raw=true)

<a name="screenshots"></a>

## Screenshots
<p>
  <p>Gambar 1. Landing Page</p>
  <img src="/img/SS1.png/">
  <nl>
  <p>Gambar 2. Visualisasi Graf Masukan (file)</p>
  <img src="/img/SS2.png/">
  <nl>
  <p>Gambar 3. Visualisasi Bridge (file)</p>
  <img src="/img/SS3.png/">
  <nl>
   <p>Gambar 4. Visualisasi SCC (file)</p>
   <img src="/img/SS4.png/">
   <nl>
   <p>Gambar 5. Visualisasi Graf (text)</p>
   <img src="/img/SS5.png/">
   <nl>
</p>

<a name="structure"></a>

## Structure
```bash
├───backend
│   └───uploads
├───frontend
│   ├───node_modules
│   ├───public
│   └───src
└───img
```

<a name="project-status">

## Project Status
Project is: _complete_

<a name="room-for-improvement">

## Room for Improvement
Perbaikan yang dapat dilakukan pada program ini adalah:
- Meningkatkan efektivitas algoritma serta menambahkan fitur-fitur lainnya

<a name="acknowledgements">

## Acknowledgements
- Terima kasih kepada Tuhan Yang Maha Esa

<a name="contact"></a>

## Contact
<h4 align="center">
  Kontak Saya : mrifki193@gmail.com<br/>
  2023
</h4>
<hr>
