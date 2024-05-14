
export default async function search(query, level) {

    try {
        return [
            { id: 1, name: query, selected: false },
            { id: 2, name: level, selected: false },
            { id: 3, name: 'File 3', selected: false }
        ]
    }
    catch (error) {
        console.error(error);
    }
}
