package main

import (
	"log"
	"net/http"
)

// Обработчик главной страницы
func home(w http.ResponseWriter, r *http.Request) {
     // Проверяется, если текущий путь URL запроса точно совпадает с шаблоном "/". Если нет, вызывается
    // функция http.NotFound() для возвращения клиенту ошибки 404.
    // Важно, чтобы мы завершили работу обработчика через return. Если мы забудем про "return", то обработчик
    // продолжит работу и выведет сообщение "Привет из SnippetBox" как ни в чем не бывало.
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

	w.Write([]byte("Привет из Snippetbox"))
}

// Обработчик для отображения содержимого заметки.
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Отображение заметки..."))
}
 
// Обработчик для создания новой заметки.
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Используем r.Method для проверки, использует ли запрос метод POST или нет. Обратите внимание,
    // что http.MethodPost является строкой и содержит текст "POST".
    if r.Method != http.MethodPost {
        // Используем метод Header().Set() для добавления заголовка 'Allow: POST' в
        // карту HTTP-заголовков. Первый параметр - название заголовка, а
        // второй параметр - значение заголовка.
        w.Header().Set("Allow", http.MethodPost)
 
        // Используем функцию http.Error() для отправки кода состояния 405 с соответствующим сообщением.
        http.Error(w, "Метод запрещен!", 405)
 
        // Затем мы завершаем работу функции вызвав "return", чтобы
        // последующий код не выполнялся.
        return
    }
	
	w.Write([]byte("Форма для создания новой заметки..."))
}

func main() {
	// Используется функция http.NewServeMux() для инициализации нового рутера, затем
	// функцию "home" регистрируется как обработчик для URL-шаблона "/".
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet

	// Используется функция http.ListenAndServe() для запуска нового веб-сервера.
	// Мы передаем два параметра: TCP-адрес сети для прослушивания (в данном случае это "localhost:4000")
	// и созданный рутер. Если вызов http.ListenAndServe() возвращает ошибку
	// мы используем функцию log.Fatal() для логирования ошибок. Обратите внимание
	// что любая ошибка, возвращаемая от http.ListenAndServe(), всегда non-nil.
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
