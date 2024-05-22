
import axiosInstance from "../configs"

export default async function getUploadedFiles() {
    try {
        let res = await axiosInstance.get(`/upload`)
        
        if (res.data.meta.code === 200) {
            res = res.data.data
            return res
        } else {
            console.error("Error fetching uploaded files: ", res.data.meta.message)
            return null
        }
    } catch (err) {
        if (err.response) {
            console.error("Get Uploaded Files Error: ", err.response.data.message)
        } else {
            console.error("Get Uploaded Files Error: ", err.message)
        }
        return null
    }
}