function addNewTask() {
    // Update displayed number
    let added_number = +document.getElementById('number-of-new-tasks').value + 1;
    document.getElementById('number-of-new-tasks').value = added_number;

    // Determine the id number of an added new task
    let id_number = 0;
    for (let i = 1; i <= added_number; i++) {
        if (document.getElementById('new-task-' + i) == null) {
            id_number = i;
            break;
        }
    }

    // Adding parts
    let new_div = document.createElement('div');
    new_div.id = 'new-task-' + id_number;
    document.getElementById('new-task-editor').appendChild(new_div);

    let new_input = document.createElement('input');
    new_input.type = 'date';
    new_input.name = 'deadline-' + id_number;
    new_input.required = true;
    document.getElementById('new-task-' + id_number).appendChild(new_input);

    let new_title = document.createElement('input');
    new_title.type = 'text';
    new_title.name = 'title-' + id_number;
    new_title.required = true;
    document.getElementById('new-task-' + id_number).appendChild(new_title);

    let new_textarea = document.createElement('textarea');
    new_textarea.name = 'new-task-' + id_number;
    new_textarea.cols = 50;
    new_textarea.rows = 10;
    new_textarea.required = true;
    document.getElementById('new-task-' + id_number).appendChild(new_textarea);

    let new_button = document.createElement('button');
    new_button.type = 'button';
    new_button.name = 'new-task-' + id_number;
    new_button.onclick = function() {reduceNewTask(id_number)};
    new_button.class = 'reduce-button';
    new_button.innerHTML = 'Reduce a new task';
    document.getElementById('new-task-' + id_number).appendChild(new_button);
}

function reduceNewTask(id_number) {
    let prev_number = document.getElementById('number-of-new-tasks').value;
    if (prev_number != 0) {
        // Update number
        document.getElementById('number-of-new-tasks').value = +prev_number - 1;
        
        // Erase the element
        document.getElementById('new-task-' + id_number).remove();
    }
}