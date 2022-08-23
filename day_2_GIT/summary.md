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

## 5 Git Commit

Perintah ini mengambil _snapshots_ dari perubahan projek saat ini

#### command

```
git commit <options>
```

#### contoh

```
git commit -a
git commit -m "your commit comment"
```

## 6. Git Push

Untuk mengirim perubahan menuju _remote repository_ (github, gitlab, dll) untuk kolaborasi

#### command

```
git push <option>
```

#### contoh

```
git push origin -u master
git push origin
```

## 7 Git Fetch

Digunakan untuk mengambil metadata terbaru dari remote ke local git, namun belum mengirimkan data apapun. Hanya memeriksa perubahan apa saja yang telah terjadi.

#### command

```
git fetch <remote> [options]
```

#### contoh

```
git fetch --all
git fetch origin
```

## 8. Git Pull

Serupa dengan perintah git pull namun kali ini file akan didownload dari remote repository dan langsung memperbarui file pada local repository.

#### command

```
git pull <remote> <branch_name>
```

#### contoh

```
git pull
git pull origin new_feature
git pull <remote repo>
```

## 9. Git branch

Untuk mengelola branch kita saat ini

#### command

```
git branch <branch>
```

#### contoh

```
git branch crazy-experiment
```

## 10. Git merge

Git merge akan menggabungkan banyak bagian dari commit menjadi satu. Umumnya digunakan untuk mengabungkan dua _branches_

#### command

```
git merge <name-branch>
```

#### contoh

```
git merge new-feature
```

## 11. Git stash

Digunakan untuk menyimapan sementara perubahan yang ada (WIP) sehingga kita bisa mengerjakan hal lain.

#### command

```
git stash <command>
```

#### perintah bisa berupa `list`, `show`, `drop`, `pop`, `apply`, `clear` dan lainnya.

#### contoh

```
git stash list
git stash apply <commit hash>
```
