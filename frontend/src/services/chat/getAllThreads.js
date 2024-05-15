
import axiosInstance from "../configs";

export default async function getAllThreads() {
    try {
        let res = await axiosInstance.get(`/thread`)
        
        res = res.data.meta.code === 200 ? res.data.data : { error: res.data.meta.message }

        return res;
    }
    catch (err) {
        if (err.response) {
            console.log(err.response.data.message);
            return { error: err.response.data.message };
        }
    }
}