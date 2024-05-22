
import axiosInstance from "../configs";

export default async function getThreadResumes(threadId) {
    try {
        let res = await axiosInstance.get(`/thread/${threadId}`)

        if (res.data.meta.code === 200) {
            res = res.data.data;
            return res;
        }
        else {
            console.error("Error fetching thread resumes: ", res.data.meta.message);
            return null;
        }
    } catch (err) {
        if (err.response) {
            console.error("Get Thread Resumes Error: ", err.response.data.message);
        } else {
            console.error("Get Thread Resumes Error: ", err.message);
        }
        return null;
    }
}
