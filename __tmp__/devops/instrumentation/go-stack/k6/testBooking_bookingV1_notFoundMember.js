import http from 'k6/http';
import { check } from 'k6';

export let options = {
  vus: 1,
  iterations: 1
};

export default function () {
  const url = `http://localhost:8080/api/v1/booking/course`;

  const payload = JSON.stringify({
    member_code: 'MEMBER-1003',
    course_code: 'COURSE-005',
  });

  const params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  const res = http.post(url, payload, params);

  check(res, {
    'status is 404': (r) => r.status === 404
  });

  try {
    const jsonData = JSON.parse(res.body);
    console.log('Response body (formatted JSON):\n', JSON.stringify(jsonData, null, 2));
  } catch (e) {
    console.log('Response is not valid JSON:', res.body);
  }
}