const apiUrl = 'http://localhost:3000/api/users';

const appElement = document.getElementById('app');

function getUsers() {
  fetch(apiUrl)
    .then(response => response.json())
    .then(data => {
      const userList = document.createElement('ul');
      data.forEach(user => {
        const first_name = document.createElement('li');
        first_name.textContent = `First name: ${user.first_name}`;
        userList.appendChild(first_name);
        const last_name = document.createElement('li');
        last_name.textContent = `Last name: ${user.last_name}`;
        userList.appendChild(last_name);
      });
      appElement.appendChild(userList);
    })
    .catch(error => {
      console.error('Error fetching data:', error);
    });
}

getUsers();
