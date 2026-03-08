import http from 'k6/http';
import { check } from 'k6';

const BASE_URL = 'http://localhost:8081/api/v1/members';

export default function () {
  let cleanDummyResponse = http.del(`${BASE_URL}/clean-dummy`);
  check(cleanDummyResponse, {
    'clean-dummy status is 200': (r) => r.status === 200,
  });

  let initDummyResponse = http.post(`${BASE_URL}/init-dummy`);
  check(initDummyResponse, {
    'init-dummy status is 200': (r) => r.status === 200,
  });

  let members = JSON.parse(initDummyResponse.body);
  let memberCodes = members.map((member) => member.code);

  return { memberCodes };
}
