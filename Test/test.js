import http from 'k6/http';
import { Counter } from 'k6/metrics';

const joinSuccess = new Counter('join_success');
const paySuccess = new Counter('pay_success');

let headers = {
    'Content-Type': 'application/json',
    'BuyerToken': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoi5bCP5piOIiwicGFzc3dkIjoiMTQ3ODUyNiJ9._zUlp5QzunD3k1xpl3BSIYHPlJNkbA_beoHK8VG-Fq4',
};

export default function () {
    let response = http.post("http://localhost:8080/buyer/join", JSON.stringify({}), { headers: headers });
    if (response.status === 200) {
        joinSuccess.add(1);
    }

    let payResponse = http.post("http://localhost:8080/buyer/pay", JSON.stringify({}), { headers: headers });
    if (payResponse.status === 200) {
        paySuccess.add(1);
    }
}