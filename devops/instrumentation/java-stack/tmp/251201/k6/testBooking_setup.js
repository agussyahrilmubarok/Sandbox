import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  vus: 1,        
  iterations: 1,
};

export default function () {
  // --- Step 1: Clean members ---
  console.log('🧹 Cleaning dummy members...');
  let res = http.del('http://localhost:8081/api/v1/members/clean-dummy');
  check(res, { 'members cleaned (200)': (r) => r.status === 200 });
  printResponse(res);

  sleep(1);

  // --- Step 2: Clean courses ---
  console.log('🧹 Cleaning dummy courses...');
  res = http.del('http://localhost:8082/api/v1/courses/clean-dummy');
  check(res, { 'courses cleaned (200)': (r) => r.status === 200 });
  printResponse(res);

  sleep(1);

  // --- Step 3: Init dummy members ---
  console.log('✨ Initializing dummy members...');
  res = http.post('http://localhost:8081/api/v1/members/init-dummy', null, {
    headers: { 'Content-Type': 'application/json' },
  });
  check(res, { 'members initialized (200)': (r) => r.status === 200 });
  printResponse(res);

  sleep(1);

  // --- Step 4: Init dummy courses ---
  console.log('✨ Initializing dummy courses...');
  res = http.post('http://localhost:8082/api/v1/courses/init-dummy', null, {
    headers: { 'Content-Type': 'application/json' },
  });
  check(res, { 'courses initialized (200)': (r) => r.status === 200 });
  printResponse(res);
}

// --- Utility: Print JSON or raw response ---
function printResponse(res) {
  try {
    const jsonData = JSON.parse(res.body);
    console.log('Response body (formatted JSON):\n', JSON.stringify(jsonData, null, 2));
  } catch (e) {
    console.log('Response is not valid JSON:', res.body);
  }
}
