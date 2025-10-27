import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:8081/api/v1/members';

export const options = {
  vus: 10,
  duration: '30s',
};

export default function () {
  let allMembersResponse = http.get(`${BASE_URL}`);
  check(allMembersResponse, {
    'find all members status is 200': (r) => r.status === 200,
  });

  sleep(1);
}
