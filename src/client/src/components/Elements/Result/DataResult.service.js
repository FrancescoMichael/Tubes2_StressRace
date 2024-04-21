import axios from 'axios';

export const getData = (callback) => {
    axios
    .get("http://localhost:8080/api/result")
    .then((res) => {
        callback(res.data)
    })
    .catch((err) => {
        console.log(err);
    });
}