
import axiosInstance from "../configs";

export default async function startThread(resumes) {
    try {

        let timeStr = new Date().toISOString().replace(/T/, ' ').replace(/\..+/, '');

        let res = await axiosInstance.post(`/thread/start`, {
            ids: resumes,
            threadName: timeStr
        });

        if (res.data.meta.code === 200) {
            return res.data.data;
        } else {
            console.error("Start Thread Error: ", res.data.meta.message);
            return null;
        }
    }
    catch (err) {
        if (err.response) {
            console.error("Start Thread Error: ", err.response.data.message);
        } else {
            console.error("Start Thread Error: ", err.message);
        }
        return null;
    }
}
