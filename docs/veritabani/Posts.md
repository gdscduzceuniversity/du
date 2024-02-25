# Posts collection'u
Bu koleksiyonda postlarla alakalı her türlü bilgiler yer alacaktır.

## _id
id alanı tipi objectid olup mongo'nun bize ön tanımlı vermiş olduğu numaralandırma için otomatik olarak kullanılmaktadır.

## author
bu alan string tipinde olup postun yazarının id'sini barındırır.

## community
bu alan string tipinde olup postun hangi topluluğa ait olduğunu barındırır.

## content
bu alan string tipinde olup postun içeriğini barındırır.

## iPaths
bu alan array tipinde olup postun içerdiği resimlerin dosya sisteminde ki adreslerini barındırır.

## createdAt
bu alan tarih tipinde olup postun oluşturulma tarihini barındırır.

## likes
bu alan array tipinde olup postu beğenen kullanıcıların id'lerini barındırır.

## comments
bu alan array tipinde olup posta yapılan yorumların id'lerini barındırır.

## tags
bu alan array tipinde olup postun etiketlerini barındırır.

## isDeleted
bu alan bool tipinde olup postun silinip silinmediğini barındırır. 1: silindi, 0: silinmedi
```