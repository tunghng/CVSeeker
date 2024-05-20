
import axiosInstance from '../configs'

export default async function uploadPdfFiles(textFiles) {

    try {
        let res = await axiosInstance.post(`/batch/upload`, {
            resumes: [...textFiles]
        });

        res = res.data.meta.code === 200 ? res.data.data : { error: res.data.meta.message }
        
        return res
    }

    catch (err) {
        if (err.response) {
            console.log(err.response.data.message);
            return { error: err.response.data.message };
        }
    }
}