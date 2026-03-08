    import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  vus: 1,
  iterations: 1,
};

export default function () {
  const headers = { 'Content-Type': 'application/json' };

  const payloadV1 = JSON.stringify({
    member_code: 'MEMBER-1000',
    course_code: 'COURSE-001',
  });

  const startV1 = Date.now();
  const resV1 = http.post('http://localhost:8080/api/v1/booking/course', payloadV1, { headers });
  const durationV1 = Date.now() - startV1;

  check(resV1, { 'v1 status is 200': (r) => r.status === 200 });
  console.log(`v1 duration: ${durationV1} ms`);
  printResponse(resV1);

  sleep(1);

  const payloadV2 = JSON.stringify({
    member_code: 'MEMBER-1001',
    course_code: 'COURSE-001',
  });

  const startV2 = Date.now();
  const resV2 = http.post('http://localhost:8080/api/v2/booking/course', payloadV2, { headers });
  const durationV2 = Date.now() - startV2;

  check(resV2, { 'v2 status is 200': (r) => r.status === 200 });
  console.log(`v2 duration: ${durationV2} ms`);
  printResponse(resV2);

  if (durationV1 < durationV2) {
    console.log(`API v1 is faster (${durationV1} ms < ${durationV2} ms)`);
  } else if (durationV1 > durationV2) {
    console.log(`API v2 is faster (${durationV2} ms < ${durationV1} ms)`);
  } else {
    console.log(`Both APIs have the same speed (${durationV1} ms)`);
  }
}

function printResponse(res) {
  try {
    const jsonData = JSON.parse(res.body);
    console.log('Response (JSON):', JSON.stringify(jsonData, null, 2));
  } catch {
    console.log('Response (raw):', res.body);
  }
}
