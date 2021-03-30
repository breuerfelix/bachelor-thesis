from os import walk, path, system, mkdir

SNIPPETS = 'snippets/'
GENERATED = 'gen/'
MINTED = 'minted.tex'
OPTIONS = 'frame=single'

if __name__ == '__main__':
    if not path.isdir(GENERATED):
        mkdir(GENERATED)

    _, _, filenames = next(walk(SNIPPETS))
    for file in filenames:
        if file == MINTED: continue

        # example command
        # pygmentize -O full -f tex -o gen/hallo.tex snippets/hallo.rs
        outFile = GENERATED + path.splitext(file)[0] + '.tex'
        system(f'pygmentize -O linenos=True -O verboptions={OPTIONS} -f tex -o {outFile} {SNIPPETS}{file}')
