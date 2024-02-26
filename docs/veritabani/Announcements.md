# Announcements collection'u
Bu koleksiyonda duyurularla alakalı her türlü bilgiler yer alacaktır.

## _id
id alanı tipi objectid olup mongo'nun bize ön tanımlı vermiş olduğu numaralandırma için otomatik olarak kullanılmaktadır.

## title
bu alan string tipinde olup duyurunun başlığını barındırır, bu alan değiştirilebilir.

## author
bu alan string tipinde olup duyurunun yazarının id'sini barındırır. Ancak bir duyurunun yazarı ancak coreTeam'den biri olabilir.

## community
bu alan string tipinde bir id olup duyurunun hangi topluluğa ait olduğunu barındırır.

## content
bu alan string tipinde olup duyurunun içeriğini barındırır.

## iPaths
bu alan array tipinde olup duyurunun içerdiği resimlerin dosya sisteminde ki adreslerini barındırır.

## createdAt
bu alan tarih tipinde olup duyurunun oluşturulma tarihini barındırır.

## likes
bu alan array tipinde olup duyuruyu beğenen kullanıcıların id'lerini barındırır.

## tags
bu alan array tipinde olup duyurunun etiketlerini barındırır.

## isDeleted
bu alan bool tipinde olup duyurunun silinip silinmediğini barındırır. 1: silindi, 0: silinmedi
