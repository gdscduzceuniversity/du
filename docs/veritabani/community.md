# Communities collection'u
bu koleksiyonda topuluklarla alakalı her türlü bilgiler yer alacaktır.

## _id
id alanı tipi objectid olup mongo'nun bize ön tanımlı vermiş olduğu numaralandırma için otomatik olarak kullanılmaktadır.

## email
bu alan string tipinde olup topluluğun emailini barındırır unique olması mühimdir.

## name
bu alan string tipinde olup topluluğun adını barındırır, bu alan değiştirilebilir.

## description
bu alan string tipinde olup topluluğun açıklamasını barındırır, bu alan değiştirilebilir.

## slug
bu alan string tipinde olup topluluğun sayfasına giden url path'ini verir, bu alan değiştirilebilir lakin unique'dir.

## iPath
bu alan string tipinde olup topluluk resminin dosya sisteminde ki adresi verilmiştir.

## members
bu alan array tipinde olup topluluğun üyelerinin id'lerini barındırır.

## memberCount
bu alan int tipinde olup topluluğun üye sayısını barındırır.

## categories
bu alan int tipinde bir array olup topluluğun kategorilerini barındırır. Kategorilerin neler olduğu frontende tutulacaktır.

## posts
bu alan array tipinde olup topluluğun postlarının id'lerini barındırır.

## postsCount
bu alan int tipinde olup topluluğun post sayısını barındırır.

## createdAt
bu alan tarih tipinde olup topluluğun oluşturulma tarihini barındırır.

## isDeleted
bu alan bool tipinde olup topluluğun silinip silinmediğini barındırır. 1: silindi, 0: silinmedi
```