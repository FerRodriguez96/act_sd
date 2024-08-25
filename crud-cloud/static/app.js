const apiUrl = 'http:///tareas';

document.addEventListener('DOMContentLoaded', obtenerTareas);

function obtenerTareas() {
    fetch(apiUrl)
        .then(response => response.json())
        .then(data => {
            const listaTareas = document.getElementById('listaTareas');
            listaTareas.innerHTML = '';
            data.forEach(tarea => {
                const li = document.createElement('li');
                li.textContent = `${tarea.nombre} `;
                li.appendChild(crearBotonEliminar(tarea.id));
                listaTareas.appendChild(li);
            });
        });
}

function crearTarea() {
    const nombre = document.getElementById('tareaInput').value;
    fetch(apiUrl, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ nombre })
    })
    .then(response => response.json())
    .then(data => {
        document.getElementById('tareaInput').value = '';
        obtenerTareas();
    });
}

function eliminarTarea(id) {
    fetch(`${apiUrl}/${id}`, {
        method: 'DELETE'
    })
    .then(response => response.json())
    .then(data => {
        obtenerTareas();
    });
}

function crearBotonEliminar(id) {
    const button = document.createElement('button');
    button.classList.add('button', 'is-danger', 'is-small');
    button.textContent = 'Eliminar';
    button.onclick = () => eliminarTarea(id);
    return button;
}