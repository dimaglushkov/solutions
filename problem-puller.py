import os
import argparse
import requests

from bs4 import BeautifulSoup


PROBLEM_URL_TEMPLATE = 'https://codeforces.com/problemset/problem/{}/{}'
PROBLEM_LOCALE = 'ru'
REPO_DIR = '/home/allacee/Repos/codeforces-solutions/'
TEST_CASES_DIR = os.path.join(REPO_DIR, 'test-cases')
LANGUAGES = {
    'python': ['.py', '#', 110]
}


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('-p', dest='pid', required=True, help='CodeWars problem id')
    parser.add_argument('-l', dest='language', required=False, default='python', help='Solution language')
    arguments = parser.parse_args()
    pid = arguments.pid
    ses = requests.session()
    ses.headers['Accept-Language'] = f'{PROBLEM_LOCALE};q=0.7,en;q=0.3'
    problem_page = ses.get(PROBLEM_URL_TEMPLATE.format(pid[:-1], pid[-1]))
    soup = BeautifulSoup(problem_page.text, 'html.parser')

    problem_desc = soup.find('div', class_='problem-statement')
    problem_title = problem_desc.find('div', class_='title').text
    problem_text = problem_desc.find_all('div', recursive=False)[1].text
    inputs = problem_desc.find('div', class_='sample-test').find_all('div', class_='input')
    outputs = problem_desc.find('div', class_='sample-test').find_all('div', class_='output')
    generate_files(pid, problem_title, problem_text, inputs, outputs)


def generate_files(pid, name, description, inputs, outputs, language='python'):
    if not os.path.isdir(TEST_CASES_DIR):
        os.mkdir(TEST_CASES_DIR)
    with open(os.path.join(TEST_CASES_DIR, pid + '.csv'), 'w') as test_case_file:
        test_case_file.write('input,output\n')
        for input_val, output_val in zip(inputs, outputs):
            test_case_file.write(f'{input_val.find("pre").text},{output_val.find("pre").text}\n')

    solution_extension = LANGUAGES[language][0]
    solution_comments = LANGUAGES[language][1]
    solution_comments_size = LANGUAGES[language][2]
    if not os.path.exists(os.path.join(REPO_DIR, pid + solution_extension)):
        with open(os.path.join(REPO_DIR, pid + solution_extension), 'w') as solution_file:
            desc_rows = int(len(description) / solution_comments_size) + 1
            comment = f'{solution_comments} {pid} {name}\n\n'
            for i in range(desc_rows):
                comment += f'{solution_comments} {description[i*solution_comments_size:(i+1)*solution_comments_size]}\n'
            solution_file.write(comment + '\n\n')


if __name__ == '__main__':
    main()
