<h1 align="center">
  Aplikasi permainan dadu
</h1>
<p align="center">
  Permainan untuk mencari point terbesar dari banyaknya dadu angka 6 yang didapat
</p>

<br/>
<br/>

# Cara menjalankan aplikasi
ketik di terminal11
```bash
go run main.go
```
kemudian 
```bash
Masukkan jumlah pemain: 5
Masukkan jumlah dadu: 5
```
dan aplikasi akan berjalan sesuai input yang diberikan

# Flow Aplikasi
### - Menentukan jumlah dadu dan jumlah pemain yang ikut permainan
### - Setiap pemain melemparkan dadunya yang dimiliki masing masing
### - Setelah itu semua pemain melakukan evaluasi
    - jika mendapat angka 6 maka pemain tersebut mendapatkan point dan dadunya dikurangi sejumlah angka 6 yang didapat
    - jika mendapat angka 1 maka dadu tersebut dioper ke pemain setelahnya
    - selain dadu angka 1 dan angka 6 dadu tersebut tetap dimainkan

### - Jika pemain sudah tidak memiliki dadu maka pemain tersebut tidak bermain lagi
### - Jika pemain tinggal satu maka permainan selesai
### - Pemenang ditentukan dari perolehan point terbesar