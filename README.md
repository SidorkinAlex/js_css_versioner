# SetterCssJS substitution of versions for css and js files in html

The program is designed to automatically update the versions of the included JavaScript (JS) and Cascading Style Sheets (CSS) files in an HTML file. This can be useful for cache-busting, ensuring that clients always load the latest versions of these files. The program takes an HTML file as input and updates the references to the JS and CSS files with a version parameter based on the current timestamp.

If you use js package builders, then it is better to configure the versioning of js and css files inside the build process

## Running the Program
To run the parser, use the following command:
```
./js_css_versioner.0.0.1  -f ./../path_to_file/index.html
```
or
```
./js_css_versioner.0.0.1  -f /path_to_file_from_cernal/index.html -o /path_to_file_from_cernal_to_write/index.html
```
Params

-f required is path to file resources html code

-o optional is path to write result. If this parameters not sett file in -f key it will be overwritten 

This programm replased your index.html file an setted version to including js and css files

before

```
<link href="css/owl.carousel.css" rel="stylesheet">
 <script src="js/jquery.js"></script>
```

after:


```
<link href="css/owl.carousel.css?v=1707405467" rel="stylesheet">
 <script src="js/jquery.js?v=1707405467"></script>
```
The program first parses the command-line parameters to obtain the file name. It then reads the content of the specified HTML file. After that, it generates a version based on the current timestamp and constructs new references for the JS and CSS files with the version parameter. Finally, it rewrites the HTML file with the updated content, including the new versioned references.

# SetterCssJS подстановка версий для css и js файлов в html

Программа предназначена для автоматического обновления версий включенных файлов JavaScript (JS) и каскадных таблиц стилей (CSS) в HTML-файле. Это может быть полезно для очистки кэша, гарантируя, что клиенты всегда загружают последние версии этих файлов. Программа принимает HTML-файл в качестве входных данных и обновляет ссылки на файлы JS и CSS параметром версии на основе текущей временной метки.

Если вы используете сборщики пакетов на js то версионирование файлов js и css лучше настроить внутри процесса сборки

## Запуск программы
Чтобы запустить, используйте следующую команду:
```
./js_css_versioner.0.0.1 -f ./../path_to_file/index.html
```
или
```
./js_css_versioner.0.0.1 -f /path_to_file_from_cernal/index.html -o /path_to_file_from_cernal_to_write/index.html
```


Параметры

-f обязательно - путь к html-коду файловых ресурсов

-o необязательно - путь для записи результата. Если этот параметр не задан в ключе -f файл будет перезаписан

Эта программа скопировала ваш файл index.html с установленной версией, включающей файлы js и css

до

```
<ссылка href="css/owl.carousel.css" rel="таблица стилей">
<script src="js/jquery.js"></script>
```

после 


```
<link href="css/owl.carousel.css?v=1707405467" rel="stylesheet">
 <script src="js/jquery.js?v=1707405467"></script>
```