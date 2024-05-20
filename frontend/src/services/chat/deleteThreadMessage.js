
import axiosInstance from "../configs";

export default async function deleteThreadMessage(threadId) {
    try {
        let res = await axiosInstance.delete(`/thread/${threadId}`);

        if (res.data.meta.code === 200) {
            return res.data.data;
        } else {
            console.error("Delete Thread Message Error: ", res.data.meta.message);
            return null;
        }
    }
    catch (err) {
        if (err.response) {
            console.error("Delete Thread Message Error: ", err.response.data.message);
        } else {
            console.error("Delete Thread Message Error: ", err.message);
        }
        return null;
    }
}