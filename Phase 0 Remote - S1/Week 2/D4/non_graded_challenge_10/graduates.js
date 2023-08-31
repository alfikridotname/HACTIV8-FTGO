function graduates(students) {
    let result = {};
    for (let i = 0; i < students.length; i++) {
        if (result[ students[ i ].class ] === undefined) {
            result[ students[ i ].class ] = [];
        }
        if (students[ i ].score > 75) {
            result[ students[ i ].class ].push({
                name: students[ i ].name,
                score: students[ i ].score
            });
        }
    }
    return result;
}

let students = [
    {
        name: 'Dimitri',
        score: 90,
        class: 'foxes'
    },
    {
        name: 'Alexei',
        score: 85,
        class: 'wolves'
    },
    {
        name: 'Sergei',
        score: 74,
        class: 'foxes'
    },
    {
        name: 'Anastasia',
        score: 78,
        class: 'wolves'
    }
];

console.log(graduates(students));
