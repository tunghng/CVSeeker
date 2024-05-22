
import axiosInstance from "../configs";

export default async function getResume(resumeId) {
    try {
        let res = await axiosInstance.get(`/${resumeId}`);

        if (res.data.meta.code === 200) {
            res = res.data.data;
            return res;
        } else {
            console.error("Error fetching resume: ", res.data.meta.message);
            return null;
        }
    } catch (err) {
        if (err.response) {
            console.error("Get Resume Error: ", err.response.data.message);
        } else {
            console.error("Get Resume Error: ", err.message);
        }
        return null;
    }
}