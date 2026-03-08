import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:8081/api/v1/members';

export default function () {
  // DELETE /clean-dummy
  let cleanDummyResponsePre = http.del(`${BASE_URL}/clean-dummy`);
  check(cleanDummyResponsePre, {
    'clean-dummy status is 200': (r) => r.status === 200,
  });

  // POST /init-dummy
  let initDummyResponse = http.post(`${BASE_URL}/init-dummy`);
  check(initDummyResponse, {
    'init-dummy status is 200': (r) => r.status === 200,
  });
  let dummyMemberCodes = JSON.parse(initDummyResponse.body).map((member) => member.code);

  // GET /members
  let allMembersResponse = http.get(`${BASE_URL}`);
  check(allMembersResponse, {
    'find all members status is 200': (r) => r.status === 200,
  });

  // GET /members/find dengan kode dummy
  dummyMemberCodes.forEach((code) => {
    let findResponse = http.get(`${BASE_URL}/find?code=${code}`);
    check(findResponse, {
      'find member by code status is 200': (r) => r.status === 200,
    });
  });

  // DELETE /clean-dummy
  let cleanDummyResponsePost = http.del(`${BASE_URL}/clean-dummy`);
  check(cleanDummyResponsePost, {
    'clean-dummy status is 200': (r) => r.status === 200,
  });

  sleep(1);
}