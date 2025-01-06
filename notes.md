### GOBO Notlar

#### **Eklenecekler**

- **Prometheus**: İzleme ve metrik toplama için entegre edilecek.
- **Ratelimit**: API Gateway katmanına taşınacak.
- **Sentry**: Hata izleme için entegre edilecek.

#### **Auth**

- Mikroservislerde **auth mekanizması** kurulmayacak.
- Tüm auth işlemleri Gateway üzerinden çözülecek.

---

#### **Testler**

1. **Load Testler**:
   - Harici bir araçla gerçekleştirilecek.
2. **Acceptance Testler**:
   - Harici bir araçla gerçekleştirilecek.
3. **SonarQube**:
   - Pipeline’da çalıştırılacak.
4. **Package Security**:
   - Pipeline’da çalıştırılacak.
5. **Contract Testler**:
   - Servislerin içinde tutulacak.

---

#### **Error Handling**

- Geliştirilecek ve bir paket haline getirilecek.

---

#### **Logger**

- Geliştirilecek ve bir paket haline getirilecek.

---

#### **Build ve Araçlar**

- **Makefile**: Proje yönetimi için oluşturulacak (alternatif olarak `justfile` kullanılabilir).
- **go-husky**:
  - Commit mesajları için bir şablon kullanılacak.
  - Credential kontrolleri yapılacak.
  - Code coverage %85’ten düşükse işlem reddedilecek.
    > Kaynak: [Go Husky Kullanımı](https://dev.to/devnull03/get-started-with-husky-for-go-31pa)

---

#### **Pipeline Planı**

1. **Lint**: Kod kalite kontrolleri.
2. **Testler**:
   - Unit Test
   - Integration Test
   - Pact Testler
3. **Code Coverage**:
   - En az %85.
4. **Build**: Projenin derlenmesi.
5. **Package Security Check**:
   - En uygun araç bulunacak, senaryolar işlenecek.
6. **Container Security Check**:
   - Docker image güvenlik taramaları.

---
