let updateTodo = function (id, isDone) {
  $.ajax({
    url: '/graphql?query=mutation+_{updateTodo(id:"' + id + '",done:' + isDone + '){id,text,done}}'
  }).done(function (data) {
    console.log(data)
    let dataParsed = JSON.parse(data)
    let updatedTodo = dataParsed.data.updateTodo
    if (updatedTodo.done) {
      $('#' + updatedTodo.id).parent().parent().parent().addClass('todo-done')
    } else {
      $('#' + updatedTodo.id).parent().parent().parent().removeClass('todo-done')
    }
  })
}

let handleTodoList = function (object) {
  let todos = object

  if (!todos.length) {
    $('.todo-list-container').append('<p>There are no tasks for you today</p>')
    return
  } else {
    $('.todo-list-container p').remove()
  }

  $.each(todos, function (i, v) {
    let todoTemplate = $('#todoItemTemplate').html()
    let todo = todoTemplate.replace('{{todo-id}}', v.id)
    todo = todo.replace('{{todo-text}}', v.text)
    todo = todo.replace('{{todo-checked}}', (v.done ? ' checked="checked"' : ''))
    todo = todo.replace('{{todo-done}}', (v.done ? ' todo-done' : ''))

    $('.todo-list-container').append(todo)
    $('#' + v.id).click(function () {
      let id = $(this).prop('id')
      let isDone = $(this).prop('checked')
      updateTodo(id, isDone)
    })
  })
}

let loadTodos = function () {
  $.ajax({
    url: "/graphql?query={todoList{id,text,done}}"
  }).done(function (data) {
    console.log(data)
    let dataParsed = JSON.parse(data)
    handleTodoList(dataParsed.data.todoList)
  })
}

let addTodo = function (todoText) {
  if (!todoText || todoText === "") {
    alert('Please specify a task')
    return
  }

  $.ajax({
    url: '/graphql?query=mutation+_{createTodo(text:"' + todoText + '"){id,text,done}}'
  }).done(function (data) {
    console.log(data)
    let dataParsed = JSON.parse(data)
    let todoList = [dataParsed.data.createTodo]
    handleTodoList(todoList)
  })
}

$(document).ready(function () {
  $('.todo-add-form').submit(function (e) {
    e.preventDefault()
    addTodo($('.todo-add-form #task').val())
    $('.todo-add-form #task').val('')
  })

  loadTodos()
})
