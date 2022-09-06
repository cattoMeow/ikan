# Cloud Native

Arsitektur dan teknologi cloud-native adalah pendekatan untuk merancang, membangun, dan mengoperasikan beban kerja yang dibangun di cloud dan memanfaatkan sepenuhnya model komputasi cloud.

Cloud Native trail map mempunyai 10 langkah:

1. Containerization
   - Umumnya menggunakan aplikasi Docker
   - Semua apps and dependency-nya dapat dibuat kontainer, sehingga kita hanya memuat kontainer untuk memuat project kita, alih-alih menyettingnya satu persatu.
2. CI/CD (Continuous Integration and Continuous Delivery)
   - Setup CI/CD so any changes in code can be automatically built, tested, and deployed.
   - Setup automated rollouts, roll backs and testing
3. Orchestration
   - Pick an orchestration solution. Kubernetes is the market leader and we should select Certified Kubernetes Platform/solution
4. Observability & Analysis
   - Memilih solusi dari pemantauan, logging, dan tracing
5. Service Mesh
   - Menyambungkan layanan bersamaan dan memberikan ingress dari interet
6. Networking
   - Untuk networking yang lebih fleksibel, gunakan CNI-Compliant network project seperti Calico, Flannel, atau Weavenet
7. Distributed Database
   - Vitess adalah pilihan terbaik untuk menjalankan mySQL.
8. Messaging
   - Untuk performa lebih tinggi dibanding JSON-REST, pertimbangkan menggunakan gRPC
9. Container Runtime
   - Umumnya menggunakan OCI-Compliant seperti rkt dan CRI-O
10. Software Distribution
    - Jika ingin melakukan distribusi software yang aman, evaluasi Notary, sebuah implementasi dari framework terupdate

# Twelve-Factor Application

Ini menjelaskan serangkaian prinsip dan praktik yang diikuti developer untuk membuat aplikasi yang dioptimalkan untuk modern cloud environment.

1. Codebase
   - Satu basis kode terkracking pada kontrol revisi
2. Dependency
   - Deklarasi dan mengisolasi dependency secara Explicitly
3. Configuration
   - Menyimpan config di dalam environment
4. Backing Service
   - Memperlakukan backing services sebagai resource yang terikat
5. Build, Release, Run
   - Memisahkan build dan run stages secara ketat
6. Process
   - Eksekusi aplikasi sebagai satu atau lebih dari satu stateless process
7. Port Binding
   - Export services melalui port binding
8. Concurrency
   - Scale out melalui model process
9. Disposability
   - Maksimalkan ketangguhan dengan startup yang cepat dan shutdown yang baik
10. Dev/Prod Parity
    - Menjaga development, staging, dan production sebaik mungkin
11. Logging
    - Memperlakukan logs sebagai event streams
12. Admin Process
    - Menjalankan admin/manajemen tugas sebagai one-off processes
