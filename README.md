.

🧭 Project Title: File Upload & Sharing System




📌 1. What this project actually does

This project is a mini Google Drive / Dropbox backend system.
It allows users to:

-Upload large files (like videos, PDFs, images)
-Download or view them later
-Share files with others through secure time-bound links
-You will build this system using:
-Golang (Gin) → for API backend
-PostgreSQL → to store metadata (file information)
-MinIO → to store actual file data
-Docker Compose → to run everything locally (Go + Postgres + MinIO)




🎯 2. Main Goal

To learn how real-world systems handle large file uploads safely, efficiently, and securely —
without crashing the backend or exposing private data.

-You will learn how to:
-Use object storage (MinIO) for saving files
-Generate presigned URLs so users can upload/download files directly
-Store metadata in PostgreSQL
-Design clean API endpoints (handlers) for upload, confirm, list, download, and share




🧩 3. Core Concept (in simple words)

Normally, when a user uploads a file → the file passes through your backend → then gets saved.
That works for small files, but not for large ones (like 500 MB videos).

So, in this project:

-Your backend doesn’t touch the big file.
-Instead, it gives the user a special temporary upload link (presigned URL).
-The user uploads directly to MinIO using that link.
-Your backend just stores information (metadata) about that file in PostgreSQL.
-That’s how cloud systems like Google Drive, Dropbox, and YouTube handle large uploads.



| Component                 | Purpose                                                |
| ------------------------- | ------------------------------------------------------ |
| 🧠 **Go Backend (Gin)**   | Handles API requests (upload, confirm, list, download) |
| 💾 **MinIO**              | Stores the actual files (as objects in buckets)        |
| 🧾 **PostgreSQL**         | Stores file information (metadata)                     |
| 🐳 **Docker Compose**     | Runs all services together locally                     |
| 🌐 **Postman / Frontend** | Used to test or trigger uploads and downloads          |





| Concept                        | What You’llUnderstand ----------------------------------------------------------------------- |
| **Object Storage**             | How files are stored as “objects” inside “buckets” (not normal folders)|
| **MinIO**                      | How to use MinIO as a self-hosted cloud storage manager    |
| **Presigned URL**              | Temporary links that let users upload/download directly to MinIO |
| **Metadata vs Data**           | Metadata (in Postgres) vs actual file data (in MinIO)      |
| **Range Requests**             | How large files (like videos) can be streamed in chunks    |
| **Share Links**                | How to generate secure, time-limited download URLs         |
| **Clean Backend Architecture** | How to separate handlers, services, and repositories in Go |




| Column         | Purpose                                                 |
| -------------- | ------------------------------------------------------- |
| `id`           | Auto-generated file ID                                  |
| `owner_id`     | ID of the user who uploaded                             |
| `filename`     | Original file name                                      |
| `bucket_name`  | Name of the MinIO bucket (e.g. `user-files`)            |
| `object_key`   | Unique key/path of file inside the bucket               |
| `size_bytes`   | File size (auto-fetched from MinIO after upload)        |
| `content_type` | Type of file (auto-fetched from MinIO)                  |
| `status`       | Upload status (`pending_upload`, `uploaded`, `deleted`) |
| `created_at`   | When record was created                                 |
| `uploaded_at`  | When upload finished                                    |





7. Project Flow (Step by Step)


Step 1️⃣ — User requests upload

-User sends file name + metadata to backend.
-Backend creates a record in Postgres (status = pending_upload).
-Backend asks MinIO for a presigned PUT URL.
-Backend sends that URL back to the user.


Step 2️⃣ — User uploads file

-User (via Postman or frontend) uses that URL to upload the file directly to MinIO.
-File gets stored on your local disk (inside MinIO’s data folder).


Step 3️⃣ — Confirm upload

-Backend checks file in MinIO (using StatObject) to get size and content-type.
-Updates Postgres record → status = uploaded.


Step 4️⃣ — Download or share

-Backend can generate presigned GET URL for secure downloads.
-Or generate share links that expire after a set time.





| Handler                      | Description                                         |
| ---------------------------- | --------------------------------------------------- |
| `PostPresignUpload`          | Create DB record + generate presigned upload URL    |
| `PostConfirmUpload`          | Confirm upload, get size/type from MinIO, update DB |
| `GetFileList`                | Show all files uploaded by a user                   |
| `GetFileDetails`             | Show details of one file                            |
| `GetPresignDownload`         | Generate presigned GET URL for download             |
| *(Optional)* `DeleteFile`    | Delete from MinIO + mark deleted in DB              |
| *(Optional)* `PostShareFile` | Generate time-bound share link                      |
| *(Optional)* `GetSharedFile` | Validate and serve shared file                      |











| Stored in        | What it stores                             |
| ---------------- | ------------------------------------------ |
| **MinIO**        | Actual file data (video, image, pdf, etc.) |
| **PostgreSQL**   | File information (metadata)                |
| **Backend (Go)** | Logic to connect them both                 |
| **Local Disk**   | Physical place where MinIO keeps files     |







🧠 10. Key Learnings from this Project

After completing it, you will know:

✅ How cloud file systems (like S3) actually work internally.

✅ How to integrate MinIO with Go using SDKs.

✅ How to design APIs for upload/download confirmation.

✅ How to separate metadata from file storage.

✅ How to use Docker Compose for multi-service apps.

✅ How to test with Postman step by step.