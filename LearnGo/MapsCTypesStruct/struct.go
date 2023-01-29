// Изучение структур и кастомных типов. Что-то типо ООП.
package main

import "fmt"

type employee struct { // Можно написать и в main(), но тогда newStruct() не будет работать, т.к
	name   string // находится за зоной видимости/доступности. Указываем глобально.
	sex    string
	age    int
	solary int
}

func main() {
	employee1 := employee{ // Обьект образца employee. Заполнен вручную.
		name:   "Mihail",
		sex:    "M",
		age:    15,
		solary: 60000,
	}
	employee2 := employee{
		name:   "Ivan",
		sex:    "M",
		age:    30,
		solary: 150000,
	}
	employee3 := newStruct("Slava", "M", 15, 60000) // Использование конструктора newStruct()

	employee4 := newStruct("Elena", "F", 24, 100000)

	fmt.Printf("Сотрудник: %s\nПол: %s\nВозраст: %d\nЗарплата: %d\n\n", employee1.name, employee1.sex, employee1.age, employee1.solary)

	setName(&employee2, "Vanya") // Заменили имя в employee2 с помощью функции setName(), где есть указатель. & - ссылка.
	fmt.Printf("Сотрудник: %s\nПол: %s\nВозраст: %d\nЗарплата: %d\n\n", employee2.name, employee2.sex, employee2.age, employee2.solary)

	fmt.Printf("%s", employee3.getInfo()) // Использование метода getInfo()

	employee4.setName2("Ekaterina")       // Аналог setName(), но в виде метода.
	fmt.Printf("%s", employee4.getInfo()) // %s - так как использутеся ftm.Sprint() (Возвращает строку)
}

func newStruct(name, sex string, age, solary int) employee { // Это называется конструктором или инициализатором.
	return employee{ // Они собирают структуры и возвращают их обьекты.
		name:   name,
		sex:    sex,
		age:    age,
		solary: solary,
	}
}

func (e employee) getInfo() string { // Метод. (e employee), где e - ресивер (переменная/получатель). Обращение к полям происходит через ресивер.
	return fmt.Sprintf("Сотрудник: %s\nПол: %s\nВозраст: %d\nЗарплата: %d\n\n", e.name, e.sex, e.age, e.solary)
}

func setName(e *employee, name string) { // Функция, не метод! Получатель и указатель на employee(1-4) как аргументы.
	e.name = name // Так делать не стоит, методы есть методы. * - указатель, убрав указатель (поле), name не будет изменятся.
}

func (e *employee) setName2(name string) { // Метод, аналогичен функции выше (54-56).
	e.name = name
}
