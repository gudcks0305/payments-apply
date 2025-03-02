// api response 
interface ApiResponse<T> {
    data?: T;
    code: number;
    message: string;
}

export default ApiResponse;