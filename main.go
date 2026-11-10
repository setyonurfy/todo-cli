package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

type TodoList struct {
	tasks    []Task
	nextID   int
	filename string
}

type TaskManager interface {
	Add(title string) error
	Update(id int, title string) error
	List() []Task
	Complete(id int) error
	Delete(id int) error
}

func NewTodoList(filename string) (*TodoList, error) {
	t := &TodoList{filename: filename, nextID: 1}

	data, err := os.ReadFile(filename)
	if errors.Is(err, os.ErrNotExist) {
		return t, nil // file belum ada, mulai dengan daftar kosong
	}
	if err != nil {
		return nil, err
	}

	if len(strings.TrimSpace(string(data))) == 0 {
		return t, nil
	}

	if err := json.Unmarshal(data, &t.tasks); err != nil {
		return nil, err
	}

	maxID := 0
	for _, task := range t.tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	t.nextID = maxID + 1

	return t, nil
}

// save menulis seluruh slice ke file JSON (method internal, tidak masuk interface)
func (t *TodoList) save() error {
	data, err := json.MarshalIndent(t.tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(t.filename, data, 0644)
}

func (t *TodoList) Add(title string) error {
	title = strings.TrimSpace(title)
	if title == "" {
		return fmt.Errorf("Task tidak boleh kosong")
	}
	newTask := Task{
		ID:    t.nextID,
		Title: title,
		Done:  false,
	}
	t.tasks = append(t.tasks, newTask)
	t.nextID++
	return t.save()
}
func (t *TodoList) Update(id int, title string) error {
	title = strings.TrimSpace(title)
	if title == "" {
		return fmt.Errorf("Judul tidak boleh kosong") // isi sendiri
	}

	for i := range t.tasks {
		if t.tasks[i].ID == id {
			t.tasks[i].Title = title
			return t.save()
		}
	}
	return fmt.Errorf("Task dengan id %d tidak ditemukan", id)
}

func (t *TodoList) List() []Task {
	return t.tasks
}

func (t *TodoList) Complete(id int) error {
	for i := range t.tasks {
		if t.tasks[i].ID == id {
			t.tasks[i].Done = true
			return t.save()
		}
	}
	return fmt.Errorf("Task dengan id %d tidak ditemukan", id)
}

func (t *TodoList) Delete(id int) error {
	for i := range t.tasks {
		if t.tasks[i].ID == id {
			t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
			return t.save()
		}
	}
	return fmt.Errorf("Task dengan ID %d tidak ditemukan", id)
}

func main() {
	todo, err := NewTodoList("tasks.json")
	if err != nil {
		fmt.Println("Gagal memuat data:", err)
		return
	}
	var manager TaskManager = todo

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n=== TO DO LIST ===")
		fmt.Println("Ketikan Angka!")
		fmt.Println("1.Tambah 2.Lihat 3.Selesai 4.Hapus 5.Edit 6.Keluar")
		fmt.Print("Pilih: ")
		scanner.Scan()
		pilihan := strings.TrimSpace(scanner.Text())

		switch pilihan {
		case "1":
			fmt.Print("Judul: ")
			scanner.Scan()
			if err := manager.Add(scanner.Text()); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Task ditambahkan")
			}

		case "2":
			tasks := manager.List()
			if len(tasks) == 0 {
				fmt.Println("Belum ada task")
				continue
			}
			for _, t := range manager.List() {
				status := " "
				if t.Done {
					status = "x"
				}
				fmt.Printf("[%d] [%s] %s\n", t.ID, status, t.Title)
			}
		case "3":
			fmt.Printf("ID: ")
			scanner.Scan()
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("ID harus berupa angka")
				continue
			}
			if err := manager.Complete(id); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Task ditandai selesai")
			}
		case "4":
			fmt.Printf("ID: ")
			scanner.Scan()
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("ID harus berupa angka")
				continue
			}
			if err := manager.Delete(id); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Task dihapus")
			}
		case "5":
			fmt.Print("ID: ")
			scanner.Scan()
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("ID harus berupa angka")
				continue
			}

			fmt.Print("Judul baru: ")
			scanner.Scan()
			if err := manager.Update(id, scanner.Text()); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Task diperbarui")
			}

		case "6":
			fmt.Println("Sampai jumpa")
			return
		default:
			fmt.Println("Pilihan tidak valid")

		}
	}

}
