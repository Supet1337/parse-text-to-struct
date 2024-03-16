# parse-text-to-struct
Программа для парсинга определенного типа строки в струткуру

Пример строк:
@id:22
@id:>30
@id:>=100
@id:!=100

#office
#office:1010
#office:>101010
#office:>=101011
#office:!=101010

Спец. символы
@ - смивол начала условия для поля
# - символ начала условия для тега
: - разделитель поля/тега и условия
> < >= <= != - операторы сравнения/условия
Если опартора сравнения нет это эквиваентно =

### Задача:

Распарсить строки и собрать из них структуру данных
type Cond struct {
	Type string
	Field string
	Condition string
	Value any
}
