import http from 'k6/http';
import execution from 'k6/execution'
import { check, sleep, fail } from 'k6';

export const options = {
  vus: 10,
  duration: '1s'
};

export default function() {
  const no = execution.vu.idInInstance
  const registerRequest = {
    email: `bambang${no}@gmail.com`,
    password: '123',
    firstName: `bambang-${no}`,
    lastName: "kentolet"
  }

  const registerResponse = http.post('http://127.0.0.1:8080/api/users', JSON.stringify(registerRequest), {
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    }
  });

  const checkRegister = check(registerResponse, {
    'register response status must 200': (res) => res.status === 200,
    'register response id must exists': (res) => res.json().data.id != null
  })
  
  if (!checkRegister) {
    fail(`Failed to register ${registerRequest.email}`) 
  }

  const loginRequest = {
    email: `bambang${no}@gmail.com`,
    password: '123',
  }

  const loginResponse = http.post('http://127.0.0.1:8080/api/users/_login', JSON.stringify(loginRequest), {
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    }
  })

  const checkLogin = check(loginResponse, {
    'login response status must 200': (res) => res.status === 200,
    'login response token must exists': (res) => res.json().data.token != ""
  })

  if (!checkLogin) {
    fail(`Failed to login user-${uniqueId}`) 
  }
  
  sleep(1)
}
