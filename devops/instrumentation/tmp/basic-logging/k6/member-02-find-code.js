import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:8081/api/v1/members';

export const options = {
  vus: 10,
  duration: '30s',
};

const memberCodes = ['MC-1XX', 'MC-2XX', 'MC-3XX'];

export default function () {
  memberCodes.forEach((code) => {
    let findResponse = http.get(`${BASE_URL}/find?code=${code}`);
    
    if (code === 'MC-3XX') {
      check(findResponse, {
        'find member by code MC-3XX status is 404': (r) => r.status === 404,
      });
    } else {
      check(findResponse, {
        'find member by code status is 200': (r) => r.status === 200,
      });
    }
  });

  sleep(1);
}
