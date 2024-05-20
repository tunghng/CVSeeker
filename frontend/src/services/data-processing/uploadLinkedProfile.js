
import axiosInstance from '../configs'

export default async function uploadLinkedProfile(linkedProfiles) {

    try {
        let res = await axiosInstance.post(`/batch/upload?isLinkedin=true`, {
            resumes: [...linkedProfiles]
        });

        console.log(res)
        
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