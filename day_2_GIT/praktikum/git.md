# Day 2 GIT

## 1. Membuat Repository Baru

Jika anda baru saja memulai git, langkah berikut sangat direkomendasikan.

```
git init <option>
```

Contoh:

```
git init --bare
git init --quiet
```

`bare` digunakan untuk membuat file git yang tidak tersembunyi, sedangkan `quiet` sebaliknya.

## 2. Mengubah _working directory_ menjadi _staging area_

Gunakan perintah ini jika kode yang sudah ditulis siap diunggah.

#### Untuk file satuan, gunakan:

```
git add <file>
```

#### Untuk direktori, gunakan:

```
git add <directory>
```

#### Contoh:

```
git add . #Untuk memilih semua file dalam repository (Sertakan tanda titik(.))
git add hello.py #Untuk file satuan
git add directory name/ #Untuk direktori
```

## 3 Menghubungkan file dari repository cloud ke local file.

Dapat digunakan untuk membuat, melihat, dan menghapus file di repository.

#### Perintah

```
git remote <options>
```

#### Contoh:

```
git remote -v
git remote add <remote name> <remote url>
```

## 4 Memeriksa status

Untuk mengecek status git apakah sudah masuk staging atau belum, serta memeriksa file yang tidak terlacak oleh git.

#### Perintah

```
git status
```

##
