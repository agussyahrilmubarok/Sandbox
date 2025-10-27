import http from 'k6/http';
import { check, sleep } from 'k6';

const MEMBER_BASE_URL = 'http://localhost:8081/api/v1/members';
const COURSE_BASE_URL = 'http://localhost:8082/api/v1/courses';
const BOOKING_BASE_URL_V1 = 'http://localhost:8080/api/v1/booking';
const BOOKING_BASE_URL_V2 = 'http://localhost:8080/api/v2/booking';

export function setup() {
    let cleanDummyMemberResponse = http.del(`${MEMBER_BASE_URL}/clean-dummy`);
    check(cleanDummyMemberResponse, {
        'clean-dummy status is 200': (r) => r.status === 200,
    });

    let cleanDummyCourseResponse = http.del(`${COURSE_BASE_URL}/clean-dummy`);
    check(cleanDummyCourseResponse, {
        'clean-dummy status is 200': (r) => r.status === 200,
    });

    let initDummyMemberResponse = http.post(`${MEMBER_BASE_URL}/init-dummy`);
    check(initDummyMemberResponse, {
        'init-dummy status is 200': (r) => r.status === 200,
    });

    let members = JSON.parse(initDummyMemberResponse.body);
    // let memberCodes = members.map((member) => member.code);

    let initDummyCourseResponse = http.post(`${COURSE_BASE_URL}/init-dummy`);
    check(initDummyCourseResponse, {
        'init-dummy status is 200': (r) => r.status === 200,
    });

    let courses = JSON.parse(initDummyCourseResponse.body);
    // let courseCodes = members.map((course) => course.code);

    return {
        members: members,
        courses: courses,
    };
}

export default function (data) {
    // FindMembers
    let findAllMemberResponse = http.get(`${MEMBER_BASE_URL}`);
    check(findAllMemberResponse, {
        'find all members status is 200': (r) => r.status === 200,
    });

    // FindMemberByCode
    const memberCodes = ['MC-1XX', 'MC-2XX', 'MC-3XX'];
    memberCodes.forEach((code) => {
        let findResponse = http.get(`${MEMBER_BASE_URL}/find?code=${code}`);

        if (code === 'MC-3XX') {
            check(findResponse, {
                'find member by code MC-3XX status is 404': (r) => r.status === 404,
            });
        } else {
            check(findResponse, {
                'find member by code status is 200': (r) => r.status === 200,
            });
        }
    });

    // FindCourses
    let findAllCoursesResponse = http.get(`${COURSE_BASE_URL}`);
    check(findAllCoursesResponse, {
        'find all courses status is 200': (r) => r.status === 200,
    });

    // FindCourseByCode
    let courseCodes = data.courses.map((course) => course.code);
    courseCodes.forEach((code) => {
        let findResponse = http.get(`${COURSE_BASE_URL}/find?code=${code}`);
        check(findResponse, {
            'find course by code status is 200': (r) => r.status === 200,
        });
    });

    // BookingCourseV1
    const bookingCoursePayloadV1 = {
        member_code: 'MC-2XX',  // Use the first member code from the setup
        course_code: 'C-002',  // Use the first course code from the setup
    };
    let bookingCourseV1Response = http.post(`${BOOKING_BASE_URL_V1}/course`, JSON.stringify(bookingCoursePayloadV1), {
        headers: { 'Content-Type': 'application/json' }
    });
    check(bookingCourseV1Response, {
        'booking course V1 status is 404': (r) => r.status === 404,
    });

    // BookingCourseV2
    const bookingCoursePayloadV2 = {
        member_code: 'MC-2XX',  // Use the first member code from the setup
        course_code: 'C-002',  // Use the first course code from the setup
    };
    let bookingCourseV2Response = http.post(`${BOOKING_BASE_URL_V2}/course`, JSON.stringify(bookingCoursePayloadV1), {
        headers: { 'Content-Type': 'application/json' }
    });
    check(bookingCourseV2Response, {
        'booking course V2 status is 404': (r) => r.status === 404,
    });

    console.log(data);
}