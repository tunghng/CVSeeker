
import axiosInstance from "../configs";

export default async function searchResume(text, level, page = 0, size = 15) {

    try {
        let res = await axiosInstance.post(`/search?knnBoost=${level}&from=${page}&size=${size}`, {
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
