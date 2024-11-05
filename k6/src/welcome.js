import { sleep } from 'k6';
import http from 'k6/http';

export const options = {
  vus: 1,
  duration: '10s',
  summaryTrendStats: ['avg', 'min', 'med', 'max', 'p(90)', 'p(95)', 'p(99)']
};

export default function() {
  http.get('http://127.0.0.1:8080/');
  sleep(1);
}
