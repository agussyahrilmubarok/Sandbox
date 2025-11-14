import http from 'k6/http';
import { check } from 'k6';

export let options = {
    vus: 1,       
    iterations: 1 
};

const url = 'http://localhost:8081/api/v1/members/init-dummy';

export default function () {
    const res = http.post(url);

    check(res, {
        'status is 200': (r) => r.status === 200
    });

    try {
        const jsonData = JSON.parse(res.body);
        console.log('Response body (formatted JSON):\n', JSON.stringify(jsonData, null, 2));
    } catch (e) {
        console.log('Response is not valid JSON:', res.body);
    }
}