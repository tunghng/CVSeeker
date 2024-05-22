
import axiosInstance from "../configs";

export default async function getAllThreads() {
    try {
        let res = await axiosInstance.get(`/thread`);
        
        if (res.data.meta.code === 200) {
            res = res.data.data;
            res.sort((a, b) => b.updated_at - a.updated_at);
            return res;
        } else {
            console.error("Error fetching threads: ", res.data.meta.message);
            return null;
        }
    } catch (err) {
        if (err.response) {
            console.error("Get All Threads Error: ", err.response.data.message);
        } else {
            console.error("Get All Threads Error: ", err.message);
        }
        return null;
    }
}
