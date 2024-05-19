
import axiosInstance from "../configs";

export default async function sendThreadMessage(threadId, message) {
    try {
        let res = await axiosInstance.post(`/thread/${threadId}/send`, {
            content: message
        });

        if (res.data.meta.code === 200) {
            return res.data.data;
        } else {
            console.error("Send Thread Message Error: ", res.data.meta.message);
            return null;
        }
    }
    catch (err) {
        if (err.response) {
            console.error("Send Thread Message Error: ", err.response.data.message);
        } else {
            console.error("Send Thread Message Error: ", err.message);
        }
        return null;
    }
}
