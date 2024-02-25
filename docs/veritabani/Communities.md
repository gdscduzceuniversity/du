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

## leader
bu alan string tipinde olup topluluğun liderinin id'sini barındırır.

## coreTeam
bu alan array tipinde olup topluluğun çekirdek ekibinin id'lerini barındırır.

## members
bu alan array tipinde olup topluluğun üyelerinin id'lerini barındırır.

## posts
bu alan array tipinde olup topluluğun postlarının id'lerini barındırır.

## createdAt
bu alan tarih tipinde olup topluluğun oluşturulma tarihini barındırır.

## isActive
bu alan bool tipinde olup topluluğun inaktif ya da aktif olduğu durumunun bilgisini içerir. 1: aktif, 0: inaktif