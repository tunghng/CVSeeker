
import axiosInstance from "../configs";

export default async function renameThreadName(threadId, newName) {    
    try {
        let res = await axiosInstance.post(`/thread/${threadId}/updateName`, {
            newName: newName
        });

        if (res.data.meta.code === 200) {
            return res.data.data;
        } else {
            console.error("Rename Thread Name Error: ", res.data.meta.message);
            return null;
        }
    }
    catch (err) {
        if (err.response) {
            console.error("Rename Thread Name Error: ", err.response.data.message);
        } else {
            console.error("Rename Thread Name Error: ", err.message);
        }
        return null;
    }
}
