
import axiosInstance from "../configs";

export default async function getThreadMessage(threadId) {
    try {
        let res = await axiosInstance.get(`/thread/${threadId}/messages`);
        
        if (res.data.meta.code === 200) {
            res = res.data.data;
            res.data.reverse();

            return res;
        } else {
            console.error("Error fetching thread messages: ", res.data.meta.message);
            return null;
        }
    } catch (err) {
        if (err.response) {
            console.error("Get Thread Message Error: ", err.response.data.message);
        } else {
            console.error("Get Thread Message Error: ", err.message);
        }
        return null;
    }
}
