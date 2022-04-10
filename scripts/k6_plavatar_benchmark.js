/*
* Benchmark for plavatar using k6 (https://github.com/grafana/k6)
* Example usage: `k6 run k6_plavatar_benchmark.js --iterations 1000 --vus 10`
 */

import http from "k6/http";

export default function () {
    let seed = Math.ceil(Math.random()*10000000);
    http.get(`http://127.0.0.1:7331/smiley/512/${seed}`);
}