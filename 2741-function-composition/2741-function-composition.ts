type F = (x: number) => number;

function compose(functions: F[]): F {
    if (functions.length === 0) {
        return (x: number) => x;
    }

    return functions.reduceRight((prev, next) => {
        return (x: number) => {
            return next(prev(x))
        }
    })
};

/**i
 * const fn = compose([x => x + 1, x => 2 * x])
 * fn(4) // 9
 */