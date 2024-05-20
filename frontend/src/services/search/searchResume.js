
import axiosInstance from "../configs";

export default async function searchResume(text, level, page = 1, size = 15) {

    let from = (page - 1) * size;

    try {
        let res = await axiosInstance.post(`/search?knnBoost=${level}&from=${from}&size=${size}`, {
            content: text
        })

        if (res.data.meta.code === 200) {
            res = res.data.data;
            return res;
        } else {
            console.log("Error searching resume: ", res.data.meta.message);
            return null;
        }
    }
    catch (err) {
        if (err.response) {
            console.error("Search Resume Error: ", err.response.data.message);
        } else {
            console.error("Search Resume Error: ", err.message);
        }
        return null;
    }
}
